package day17

import (
	"container/heap"
	"errors"
 
	"github.com/ndfsa/advent-of-code-2023/util"
)

var (
	DIR_UP    = util.DIR_UP
	DIR_DOWN  = util.DIR_DOWN
	DIR_RIGHT = util.DIR_RIGHT
	DIR_LEFT  = util.DIR_LEFT
)

type StateHeap []State

func (b StateHeap) Len() int {
	return len(b)
}

func (b StateHeap) Less(i, j int) bool {
	return b[i].loss < b[j].loss
}

func (b *StateHeap) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *StateHeap) Push(element any) {
	*b = append(*b, element.(State))
}

func (b *StateHeap) Pop() any {
	length := len(*b)
	res := (*b)[length-1]
	*b = (*b)[:length-1]
	return res
}

type Block struct {
	pos   util.Vec2
	dir   util.Vec2
	count int
}

type State struct {
	block Block
	loss  int
}

func (t Block) neighbors(field [][]byte, minSteps, maxSteps int) []Block {
	next := []util.Vec2{
		DIR_UP,
		DIR_RIGHT,
		DIR_DOWN,
		DIR_LEFT,
	}

	res := []Block{}
	for _, dir := range next {
		next := t.pos.Add(dir)

		if next.Row < 0 ||
			next.Row >= len(field) ||
			next.Col < 0 ||
			next.Col >= len(field[0]) ||
			t.dir == (util.Vec2{Row: -dir.Row, Col: -dir.Col}) {
			continue
		}

		if t.dir != dir && t.count >= minSteps {
			res = append(res, Block{
				pos:   next,
				dir:   dir,
				count: 1})
		}
		if t.dir == dir && t.count < maxSteps {
			res = append(res, Block{
				pos:   next,
				dir:   dir,
				count: t.count + 1})
		}
	}
	return res
}

func parseInput(lines []string) [][]byte {
	res := [][]byte{}
	for _, line := range lines {
		row := []byte{}
		for i := range line {
			row = append(row, line[i]-'0')
		}
		res = append(res, row)
	}
	return res
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	return minPath(parseInput(lines), 0, 3)
}

func minPath(cityBlocks [][]byte, minSteps int, maxSteps int) (int, error) {
	start := util.Vec2{Row: 0, Col: 0}
	end := util.Vec2{Row: len(cityBlocks) - 1, Col: len(cityBlocks[0]) - 1}

	visited := map[Block]struct{}{}

	queue := &StateHeap{}
	heap.Init(queue)
	heap.Push(queue, State{block: Block{pos: start, dir: DIR_RIGHT}})
	heap.Push(queue, State{block: Block{pos: start, dir: DIR_DOWN}})

	for queue.Len() > 0 {
		state := heap.Pop(queue).(State)

		if state.block.pos == end && state.block.count >= minSteps {
			return state.loss, nil
		}

		if _, ok := visited[state.block]; ok {
			continue
		}
		visited[state.block] = struct{}{}

		for _, neighbor := range state.block.neighbors(cityBlocks, minSteps, maxSteps) {
			newCost := int(cityBlocks[neighbor.pos.Row][neighbor.pos.Col]) + state.loss
			heap.Push(queue, State{block: neighbor, loss: newCost})
		}
	}

	return 0, errors.New("not found")
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	return minPath(parseInput(lines), 4, 10)
}

package day17

import (
	"container/heap"
	"errors"

	"github.com/ndfsa/advent-of-code-2023/util"
)

var (
	DIR_UP    = util.DIR_V2_NEG_X
	DIR_DOWN  = util.DIR_V2_POS_X
	DIR_RIGHT = util.DIR_V2_POS_Y
	DIR_LEFT  = util.DIR_V2_NEG_Y
)

type Block struct {
	pos   util.Vec2
	dir   util.Vec2
	count int
}

func (t Block) neighbors(field [][]byte, minSteps, maxSteps int) []Block {
	neighbors := t.pos.Neighbors(func(v util.Vec2, d util.Vec2) bool {
		return v.X >= 0 &&
			v.X < len(field) &&
			v.Y >= 0 &&
			v.Y < len(field[0])
	})
	res := []Block{}
	for _, neighbor := range neighbors {

		dir := neighbor.Sus(t.pos)

		if t.dir == (util.Vec2{X: -dir.X, Y: -dir.Y}) {
			continue
		}
		if t.dir != dir && t.count >= minSteps {
			res = append(res, Block{
				pos:   neighbor,
				dir:   dir,
				count: 1})
		}
		if t.dir == dir && t.count < maxSteps {
			res = append(res, Block{
				pos:   neighbor,
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
	start := util.Vec2{X: 0, Y: 0}
	end := util.Vec2{X: len(cityBlocks) - 1, Y: len(cityBlocks[0]) - 1}

	visited := map[Block]struct{}{}

	queue := &util.StateHeap[Block]{}
	heap.Init(queue)
	heap.Push(queue, util.State[Block]{Vertex: Block{pos: start, dir: DIR_RIGHT}})
	heap.Push(queue, util.State[Block]{Vertex: Block{pos: start, dir: DIR_DOWN}})

	for queue.Len() > 0 {
		state := heap.Pop(queue).(util.State[Block])

		if state.Vertex.pos == end && state.Vertex.count >= minSteps {
			return state.Cost, nil
		}

		if _, ok := visited[state.Vertex]; ok {
			continue
		}
		visited[state.Vertex] = struct{}{}

		for _, neighbor := range state.Vertex.neighbors(cityBlocks, minSteps, maxSteps) {
			newCost := int(cityBlocks[neighbor.pos.X][neighbor.pos.Y]) + state.Cost
			heap.Push(queue, util.State[Block]{Vertex: neighbor, Cost: newCost})
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

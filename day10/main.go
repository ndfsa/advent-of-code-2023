package day10

import (
	"errors"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Blueprint struct {
	data []string
}

func (b Blueprint) getPipe(p util.Vec2) (byte, error) {
	rowMax, colMax := len(b.data), len(b.data[0])
	if p.Row < 0 || p.Row >= rowMax || p.Col < 0 || p.Col >= colMax {
		return 0, errors.New("invalid position")
	}

	return b.data[p.Row][p.Col], nil
}

func (b Blueprint) getStartAdjacent() (util.Vec2, util.Vec2, util.Vec2) {
	var pipesStart util.Vec2
	for idx, row := range b.data {
		if col := strings.IndexByte(row, 'S'); col != -1 {
			pipesStart = util.Vec2{Row: idx, Col: col}
			break
		}
	}
	startNext := b.findNext(pipesStart)
	return pipesStart, startNext[0], startNext[1]
}

func (b Blueprint) findNext(p util.Vec2) []util.Vec2 {
	currentPipe, _ := b.getPipe(p)

	possibleNext := []util.Vec2{
		{Row: p.Row - 1, Col: p.Col},
		{Row: p.Row + 1, Col: p.Col},
		{Row: p.Row, Col: p.Col - 1},
		{Row: p.Row, Col: p.Col + 1},
	}

	next := []util.Vec2{}

	for _, np := range possibleNext {
		pipe, err := b.getPipe(np)
		if err != nil {
			continue
		}
		if np.Row < p.Row && (pipe == '|' || pipe == '7' || pipe == 'F') &&
			(currentPipe == 'S' || currentPipe == '|' || currentPipe == 'J' || currentPipe == 'L') {
			next = append(next, np)
		} else if np.Row > p.Row && (pipe == '|' || pipe == 'L' || pipe == 'J') &&
			(currentPipe == 'S' || currentPipe == '|' || currentPipe == '7' || currentPipe == 'F') {
			next = append(next, np)
		} else if np.Col < p.Col && (pipe == '-' || pipe == 'L' || pipe == 'F') &&
			(currentPipe == 'S' || currentPipe == '-' || currentPipe == '7' || currentPipe == 'J') {
			next = append(next, np)
		} else if np.Col > p.Col && (pipe == '-' || pipe == '7' || pipe == 'J') &&
			(currentPipe == 'S' || currentPipe == '-' || currentPipe == 'L' || currentPipe == 'F') {
			next = append(next, np)
		}
	}
	return next
}

type State struct {
	pos      util.Vec2
	start    util.Vec2
	end      util.Vec2
	previous map[util.Vec2]util.Vec2
}

func (s *State) findLoop(b Blueprint) {
	stack := b.findNext(s.pos)
	for s.pos != s.end {
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s.previous[next] = s.pos

		s.pos = next

		possibleNext := b.findNext(s.pos)
		for _, point := range possibleNext {
			if _, ok := s.previous[point]; !ok {
				stack = append(stack, point)
			}
		}
	}

	s.previous[s.start] = s.end
}

func parseInput(lines []string) Blueprint {
	return Blueprint{data: lines}
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	b := parseInput(lines)
	pipesStart, start, end := b.getStartAdjacent()

	state := State{start: pipesStart, end: end, pos: start, previous: map[util.Vec2]util.Vec2{start: pipesStart}}
	state.findLoop(b)

	return len(state.previous) / 2, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	b := parseInput(lines)
	pipesStart, start, end := b.getStartAdjacent()

	state := State{start: pipesStart, end: end, pos: start, previous: map[util.Vec2]util.Vec2{start: pipesStart}}
	state.findLoop(b)

	area := 0
	curr := start

	// Shoelace formula
	for i := 0; i < len(state.previous); i++ {
		prev := state.previous[curr]

		area += curr.Row*prev.Col - prev.Row*curr.Col

		curr = prev
	}
	area /= 2

	// Pick's theorem
	return area - (len(state.previous) / 2) + 1, nil
}

package day08

import (
	"errors"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Point struct {
	row int
	col int
}

type Blueprint struct {
	data []string
}

func (b Blueprint) getPipe(p Point) (byte, error) {
	rowMax, colMax := len(b.data), len(b.data[0])
	if p.row < 0 || p.row >= rowMax || p.col < 0 || p.col >= colMax {
		return 0, errors.New("invalid position")
	}

	return b.data[p.row][p.col], nil
}

func (b Blueprint) getStartAdjacent() (Point, Point, Point) {
	var pipesStart Point
	for idx, row := range b.data {
		if col := strings.IndexByte(row, 'S'); col != -1 {
			pipesStart = Point{row: idx, col: col}
			break
		}
	}
	startNext := b.findNext(pipesStart)
	return pipesStart, startNext[0], startNext[1]
}

func (b Blueprint) findNext(p Point) []Point {
	currentPipe, _ := b.getPipe(p)

	possibleNext := []Point{
		{p.row - 1, p.col},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
		{p.row, p.col + 1},
	}

	next := []Point{}

	for _, np := range possibleNext {
		pipe, err := b.getPipe(np)
		if err != nil {
			continue
		}
		if np.row < p.row && (pipe == '|' || pipe == '7' || pipe == 'F') &&
			(currentPipe == 'S' || currentPipe == '|' || currentPipe == 'J' || currentPipe == 'L') {
			next = append(next, np)
		} else if np.row > p.row && (pipe == '|' || pipe == 'L' || pipe == 'J') &&
			(currentPipe == 'S' || currentPipe == '|' || currentPipe == '7' || currentPipe == 'F') {
			next = append(next, np)
		} else if np.col < p.col && (pipe == '-' || pipe == 'L' || pipe == 'F') &&
			(currentPipe == 'S' || currentPipe == '-' || currentPipe == '7' || currentPipe == 'J') {
			next = append(next, np)
		} else if np.col > p.col && (pipe == '-' || pipe == '7' || pipe == 'J') &&
			(currentPipe == 'S' || currentPipe == '-' || currentPipe == 'L' || currentPipe == 'F') {
			next = append(next, np)
		}
	}
	return next
}

type State struct {
	pos      Point
	start    Point
	end      Point
	previous map[Point]Point
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

	state := State{start: pipesStart, end: end, pos: start, previous: map[Point]Point{start: pipesStart}}
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

	state := State{start: pipesStart, end: end, pos: start, previous: map[Point]Point{start: pipesStart}}
	state.findLoop(b)

	area := 0
	curr := start

	// Shoelace formula
	for i := 0; i < len(state.previous); i++ {
		prev := state.previous[curr]

		area +=  curr.row*prev.col - prev.row*curr.col

		curr = prev
	}
	area /= 2

	// Pick's theorem
	return area - (len(state.previous) / 2) + 1, nil
}

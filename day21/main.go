package day21

import (
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Plot struct {
	pos   util.Vec2
	steps int
}

func parseInput(lines []string) ([][]byte, util.Vec2) {
	res := [][]byte{}

	row, col := 0, 0
	for i, line := range lines {
		if j := strings.IndexByte(line, 'S'); j != -1 {
			row, col = i, j
			break
		}
		res = append(res, []byte(line))
	}

	for _, line := range lines[row:] {
		res = append(res, []byte(line))
	}

	return res, util.Vec2{Row: row, Col: col}
}

func SolvePart1(dist int) util.Solution[int] {
	return util.Solution[int](func(filePath string) (int, error) {
		return solveP1(filePath, dist)
	})
}

func solveP1(filePath string, dist int) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	field, start := parseInput(lines)

	queue := []Plot{{pos: start, steps: dist}}
	visited := map[util.Vec2]struct{}{}

	res := 0
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.steps < 0 {
			continue
		}

		if _, seen := visited[next.pos]; seen {
			continue
		}
		visited[next.pos] = struct{}{}

		if next.steps%2 == 0 {
			res++
		}
		neighbors := next.pos.Neighbors(func(v util.Vec2) bool {
			return v.Row >= 0 &&
				v.Row < len(field) &&
				v.Col >= 0 &&
				v.Col < len(field[0]) &&
				field[v.Row][v.Col] != '#'
		})

		for _, v := range neighbors {
			queue = append(queue, Plot{pos: v, steps: next.steps - 1})
		}
	}

	return res, nil
}

func SolvePart2(dist int) util.Solution[int] {
	return util.Solution[int](func(filePath string) (int, error) {
		return solveP2(filePath, dist)
	})
}

func solveP2(filePath string, dist int) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	parseInput(lines)

	return 0, nil
}

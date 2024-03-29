package day21

import (
	"math"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Tile struct {
	pos    util.Vec2
	offset util.Vec2
}

type Plot struct {
	tile  Tile
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

	return res, util.Vec2{X: row, Y: col}
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

	queue := []Plot{{tile: Tile{pos: start}, steps: dist}}
	visited := map[util.Vec2]struct{}{}

	res := 0
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.steps < 0 {
			continue
		}

		if _, seen := visited[next.tile.pos]; seen {
			continue
		}
		visited[next.tile.pos] = struct{}{}

		if next.steps%2 == 0 {
			res++
		}
		neighbors := next.tile.pos.Neighbors(func(v util.Vec2, dir util.Vec2) bool {
			return v.X >= 0 &&
				v.X < len(field) &&
				v.Y >= 0 &&
				v.Y < len(field[0]) &&
				field[v.X][v.Y] != '#'
		})

		for _, v := range neighbors {
			queue = append(queue, Plot{tile: Tile{pos: v}, steps: next.steps - 1})
		}
	}

	return res, nil
}

func mod(x, m int) int {
	return (x%m + m) % m
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

	field, start := parseInput(lines)

	rows := len(field)
	cols := len(field[0])

	queue := []Plot{{tile: Tile{pos: start}, steps: dist}}
	visited := map[Tile]struct{}{}
	answer := map[Tile]struct{}{}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.steps < 0 {
			continue
		}

		if _, seen := visited[next.tile]; seen {
			continue
		}
		visited[next.tile] = struct{}{}

		if next.steps%2 == 0 {
			answer[next.tile] = struct{}{}
		}

		neighbors := next.tile.pos.Neighbors(func(v util.Vec2, dir util.Vec2) bool {
			r := mod(v.X, rows)
			c := mod(v.Y, cols)
			return field[r][c] != '#'
		})

		for _, v := range neighbors {
			tile := Tile{
				pos: util.Vec2{
					X: mod(v.X, rows),
					Y: mod(v.Y, cols),
				},
				offset: next.tile.offset.Add(
					util.Vec2{
						X: int(math.Floor(float64(v.X) / float64(rows))),
						Y: int(math.Floor(float64(v.Y) / float64(cols)))})}

			queue = append(queue, Plot{tile: tile, steps: next.steps - 1})
		}
	}

	return len(answer), nil
}

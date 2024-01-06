package day18

import (
	"fmt"
	"strconv"

	"github.com/ndfsa/advent-of-code-2023/util"
)

var (
	DIR_UP    = util.DIR_UP
	DIR_DOWN  = util.DIR_DOWN
	DIR_RIGHT = util.DIR_RIGHT
	DIR_LEFT  = util.DIR_LEFT
)

type Vertex struct {
	pos   util.Vec2
	color string
}

func parseDirection(d byte) util.Vec2 {
	switch d {
	case 'U', '3':
		return DIR_UP
	case 'D', '1':
		return DIR_DOWN
	case 'R', '0':
		return DIR_RIGHT
	case 'L', '2':
		return DIR_LEFT
	}
	panic("unknown direction")
}

func parseInput(lines []string) []Vertex {
	res := []Vertex{{}}
	for _, line := range lines {
		var dir byte
		var count int
		var color string

		fmt.Sscanf(line, "%c %d (#%s)", &dir, &count, &color)
		next := Vertex{
			pos:   res[len(res)-1].pos.AddMult(parseDirection(dir), count),
			color: color}

		res = append(res, next)
	}

	return res
}

func calculateArea(vertices []Vertex) int {
	area := 0
	boundary := 0

	// Shoelace formula again
	for i := 1; i < len(vertices); i++ {
		prev := vertices[i-1].pos
		curr := vertices[i].pos

		boundary += curr.HammingDist(prev)
		area += curr.Row*prev.Col - prev.Row*curr.Col
	}
	area /= 2

	// Pick's theorem again
	return area + 1 + boundary/2
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	vertices := parseInput(lines)

	return calculateArea(vertices), nil

}

func decodeInstructions(mangled []Vertex) []Vertex {
	res := []Vertex{{}}
	for _, vert := range mangled[1:] {
		n, _ := strconv.ParseInt(vert.color[:5], 16, 64)

		next := Vertex{pos: res[len(res)-1].pos.AddMult(parseDirection(vert.color[5]), int(n))}
		res = append(res, next)
	}

	return res
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	vertices := parseInput(lines)
	vertices = decodeInstructions(vertices)
	return calculateArea(vertices), nil
}

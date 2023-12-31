package day14

import (
	"fmt"

	"github.com/ndfsa/advent-of-code-2023/util"
)

var (
	DIR_NORTH = util.Point{Row: -1, Col: 0}
	DIR_SOUTH = util.Point{Row: 1, Col: 0}
	DIR_EAST  = util.Point{Row: 0, Col: 1}
	DIR_WEST  = util.Point{Row: 0, Col: -1}
)

type Dish struct {
	roundedRocks map[util.Point]struct{}
	cubeRocks    map[util.Point]struct{}
	height       int
	width        int
}

func (d Dish) String() string {
	res := ""
	for i := 0; i < d.height; i++ {
		for j := 0; j < d.width; j++ {
			pos := util.Point{Row: i, Col: j}
			if _, ok := d.roundedRocks[pos]; ok {
				res += "O"
				continue
			}

			if _, ok := d.cubeRocks[pos]; ok {
				res += "#"
				continue
			}
			res += "."
		}
		res += "\n"
	}

	return res
}

func (d *Dish) calculateLoad() int {
	res := 0
	for rock := range d.roundedRocks {
		res += d.height - rock.Row
	}
	return res
}

func (d *Dish) tilt(direction util.Point) {
	for {
		nextRoundedRocks := make(map[util.Point]struct{})
		moved := 0
		for rock := range d.roundedRocks {
			nextPos := rock.Sum(direction)

			if nextPos.Row < 0 ||
				nextPos.Row > d.height ||
				nextPos.Col < 0 ||
				nextPos.Col > d.width {

				nextRoundedRocks[rock] = struct{}{}
				continue
			}

			if _, ok := d.roundedRocks[nextPos]; ok {
				nextRoundedRocks[rock] = struct{}{}
				continue
			}

			if _, ok := d.cubeRocks[nextPos]; ok {
				nextRoundedRocks[rock] = struct{}{}
				continue
			}

			nextRoundedRocks[nextPos] = struct{}{}
			moved++
		}

		if moved == 0 {
			break
		}
		d.roundedRocks = nextRoundedRocks
		moved = 0
	}
}

func parseInput(lines []string) Dish {
	dish := Dish{
		roundedRocks: make(map[util.Point]struct{}),
		cubeRocks:    make(map[util.Point]struct{}),
		height:       len(lines),
		width:        len(lines[0])}
	for r, line := range lines {
		for c, ch := range line {
			switch ch {
			case 'O':
				dish.roundedRocks[util.Point{Row: r, Col: c}] = struct{}{}
			case '#':
				dish.cubeRocks[util.Point{Row: r, Col: c}] = struct{}{}
			}
		}
	}
	return dish
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	dish := parseInput(lines)
	dish.tilt(DIR_NORTH)

	return dish.calculateLoad(), nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	dish := parseInput(lines)
	dish.tilt(DIR_NORTH)
	dish.tilt(DIR_WEST)
	dish.tilt(DIR_SOUTH)
	dish.tilt(DIR_EAST)

	fmt.Printf("dish:\n%v\n", dish)

	return dish.calculateLoad(), nil
}

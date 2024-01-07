package day14

import (
	"crypto/md5"
	"fmt"

	"github.com/ndfsa/advent-of-code-2023/util"
)

var (
	DIR_NORTH = util.DIR_V2_NEG_X
	DIR_SOUTH = util.DIR_V2_POS_X
	DIR_EAST  = util.DIR_V2_POS_Y
	DIR_WEST  = util.DIR_V2_NEG_Y
)

const (
	TYPE_EMPTY byte = iota
	TYPE_ROUNDED
	TYPE_CUBED
)

type Dish struct {
	rocks  [][]byte
	height int
	width  int
}

func (d Dish) stateHash() string {
	data := []byte{}
	for _, row := range d.rocks {
		for _, b := range row {
			data = append(data, b)
		}
	}

	return fmt.Sprintf("%x", md5.Sum(data))
}

func (d *Dish) calculateLoad() int {
	res := 0
	for i := 0; i < d.height; i++ {
		for j := 0; j < d.width; j++ {
			if d.rocks[i][j] == TYPE_ROUNDED {
				res += d.height - i
			}
		}
	}
	return res
}

func (d Dish) valid(pos util.Vec2) bool {
	return pos.X >= 0 &&
		pos.X < d.height &&
		pos.Y >= 0 &&
		pos.Y < d.width
}

func (d *Dish) tilt(direction util.Vec2) {
	var i, j int
	switch direction {
	case DIR_NORTH:
		i, j = 1, 0
	case DIR_SOUTH:
		i, j = d.height-2, 0
	case DIR_EAST:
		i, j = 0, d.width-2
	case DIR_WEST:
		i, j = 0, 1
	}
	for i >= 0 && i < d.height && j >= 0 && j < d.width {

		pos := util.Vec2{X: i, Y: j}
		if d.rocks[i][j] == TYPE_ROUNDED {
			nextPos := pos.Add(direction)

			for d.valid(nextPos) && d.rocks[nextPos.X][nextPos.Y] == TYPE_EMPTY {
				nextPos = nextPos.Add(direction)
			}
			nextPos = nextPos.Sus(direction)

			d.rocks[i][j] = TYPE_EMPTY
			d.rocks[nextPos.X][nextPos.Y] = TYPE_ROUNDED
		}

		switch direction {
		case DIR_NORTH, DIR_SOUTH:
			j++
			if j == d.width {
				i -= direction.X
				j = 0
			}
		case DIR_EAST, DIR_WEST:
			i++
			if i == d.height {
				j -= direction.Y
				i = 0
			}
		}
	}
}

func parseInput(lines []string) Dish {
	dish := Dish{
		rocks:  [][]byte{},
		height: len(lines),
		width:  len(lines[0])}
	for _, line := range lines {
		row := []byte{}
		for _, ch := range line {
			switch ch {
			case '.':
				row = append(row, TYPE_EMPTY)
			case 'O':
				row = append(row, TYPE_ROUNDED)
			case '#':
				row = append(row, TYPE_CUBED)
			}
		}
		dish.rocks = append(dish.rocks, row)
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
	states := make(map[string]int)

	limit := 1_000_000_000

	for i := 0; i < limit; i++ {
		dish.tilt(DIR_NORTH)
		dish.tilt(DIR_WEST)
		dish.tilt(DIR_SOUTH)
		dish.tilt(DIR_EAST)

		hash := dish.stateHash()
		if prev, ok := states[hash]; ok {
			limit = (limit-prev)%(i-prev) + i
		} else {
			states[hash] = i
		}
	}

	return dish.calculateLoad(), nil
}

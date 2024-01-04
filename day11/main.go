package day11

import (
	"github.com/ndfsa/advent-of-code-2023/util"
)

func parseInput(lines []string, expansion int) []util.Vec2 {
	galaxies := []util.Vec2{}
	rowGravity := make([]bool, len(lines))
	colGravity := make([]bool, len(lines[0]))
	for i, line := range lines {
		for j, ch := range line {
			if ch != '.' {
				galaxies = append(galaxies, util.Vec2{Row: i, Col: j})
				rowGravity[i] = true
				colGravity[j] = true
			}
		}
	}

	offset := 0
	for i, gravity := range rowGravity {
		if gravity {
			continue
		}

		for idx := range galaxies {
			if galaxies[idx].Row-offset > i {
				galaxies[idx].Row += expansion - 1
			}
		}
		offset += expansion - 1
	}

	offset = 0
	for i, gravity := range colGravity {
		if gravity {
			continue
		}

		for idx := range galaxies {
			if galaxies[idx].Col-offset > i {
				galaxies[idx].Col += expansion - 1
			}
		}
		offset += expansion - 1
	}
	return galaxies
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	galaxies := parseInput(lines, 2)

	res := 0
	for idx, g1 := range galaxies {
		for _, g2 := range galaxies[idx+1:] {
			res += util.AbsDiff(g1.Row, g2.Row)
			res += util.AbsDiff(g1.Col, g2.Col)
		}
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	galaxies := parseInput(lines, 1_000_000)

	res := 0
	for idx, g1 := range galaxies {
		for _, g2 := range galaxies[idx+1:] {
			res += util.AbsDiff(g1.Row, g2.Row)
			res += util.AbsDiff(g1.Col, g2.Col)
		}
	}

	return res, nil
}

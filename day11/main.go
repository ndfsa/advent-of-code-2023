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
				galaxies = append(galaxies, util.Vec2{X: i, Y: j})
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
			if galaxies[idx].X-offset > i {
				galaxies[idx].X += expansion - 1
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
			if galaxies[idx].Y-offset > i {
				galaxies[idx].Y += expansion - 1
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
			res += util.AbsDiff(g1.X, g2.X)
			res += util.AbsDiff(g1.Y, g2.Y)
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
			res += util.AbsDiff(g1.X, g2.X)
			res += util.AbsDiff(g1.Y, g2.Y)
		}
	}

	return res, nil
}

package day13

import (
	"errors"
	"math/bits"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Pattern struct {
	data  []int
	width int
}

func (p *Pattern) transpose() {
	tr := make([]int, p.width)
	for i := range tr {
		offset := p.width - 1 - i
		for j := len(p.data) - 1; j >= 0; j-- {
			tr[i] |= p.data[j] & (1 << offset) >> offset << j
		}
	}
	p.width = len(p.data)
	p.data = tr
}

func (p Pattern) getSummary(smudge bool) (int, error) {
	i := 1
outer:
	for ; i < len(p.data); i++ {
		above := []int{}
		below := []int{}

		above = append(above, p.data[:i]...)
		below = append(below, p.data[i:]...)

		for j, k := 0, len(above)-1; j < k; j, k = j+1, k-1 {
			above[j], above[k] = above[k], above[j]
		}

		mn := min(len(above), len(below))
		above = above[:mn]
		below = below[:mn]

		if smudge {
			diff := 0
			for i, v := range above {
				diff += bits.OnesCount(uint(below[i] ^ v))
			}
			if diff != 1 {
				continue outer
			}
		} else if !util.SlicesEqual(above, below) {
			continue outer
		}

		return i, nil
	}
	return 0, errors.New("no mirror found")
}

func addSummaries(patterns []Pattern, smudge bool) int {
	res := 0
	for _, pattern := range patterns {
		summary, err := pattern.getSummary(smudge)
		if err != nil {
			pattern.transpose()
			summary, _ = pattern.getSummary(smudge)
			res += summary
		} else {
			res += summary * 100
		}
	}
	return res
}

func parseInput(input string) []Pattern {
	patterns := []Pattern{}
	for _, chunk := range strings.Split(input, "\n\n") {
		pattern := Pattern{data: []int{}}

		rows := strings.Split(chunk, "\n")
		pattern.width = len(rows[0])

		for _, line := range rows {
			row := 0
			for i, ch := range line {
				if ch == '#' {
					row |= 1 << (pattern.width - i - 1)
				}
			}
			pattern.data = append(pattern.data, row)
		}
		patterns = append(patterns, pattern)
	}

	return patterns
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	patterns := parseInput(input)

	return addSummaries(patterns, false), nil
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	patterns := parseInput(input)

	return addSummaries(patterns, true), nil
}

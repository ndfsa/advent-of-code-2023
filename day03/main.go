package day03

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/ndfsa/advent-of-code-2023/util"
)

const (
	OFF = 0
	ON  = 1
)

type Number struct {
	row   int
	start int
	end   int
	found bool
}

type Part struct {
	row   int
	col   int
	value rune
}

func (n Number) GetNumberValue(lines []string) int {
	part, _ := strconv.Atoi(lines[n.row][n.start:n.end])
	return part
}

func inRange(part Part, num Number) bool {
	return num.row-1 <= part.row && part.row <= num.row+1 &&
		num.start-1 <= part.col && part.col <= num.end

}

func parseInput(lines []string) ([]Part, []Number) {
	parts := make([]Part, 0)
	nums := make([]Number, 0)

	re := regexp.MustCompile(`([0-9])+`)
	for i, line := range lines {
		for j, ch := range line {
			if ch != '.' && !unicode.IsDigit(ch) {
				parts = append(parts, Part{row: i, col: j, value: ch})
			}
		}

		for _, match := range re.FindAllIndex([]byte(line), -1) {
			nums = append(nums, Number{start: match[0], end: match[1], row: i, found: false})
		}
	}
	return parts, nums
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)
	if err != nil {
		return 0, err
	}

	parts, nums := parseInput(lines)
	for _, part := range parts {
		for i, num := range nums {
			if inRange(part, num) {
				nums[i].found = true
			}
		}
	}

	res := 0
	for _, num := range nums {
		if num.found {
			res += num.GetNumberValue(lines)
		}
	}
	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)
	if err != nil {
		return 0, err
	}

	parts, nums := parseInput(lines)
	res := 0
out:
	for _, part := range parts {
		if part.value != '*' {
			continue
		}

		found := make([]Number, 0)
		for _, num := range nums {
			if inRange(part, num) {
				found = append(found, num)
			}

			if len(found) > 2 {
				continue out
			}
		}
		if len(found) != 2 {
			continue out
		}

		res += (found[0].GetNumberValue(lines) * found[1].GetNumberValue(lines))
	}

	return res, nil
}

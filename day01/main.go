package day01

import (
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	fRes := 0
	for _, line := range lines {
		pRes := 0
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if util.IsDigit(ch) {
				pRes *= 10
				pRes += int(ch - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if util.IsDigit(ch) {
				pRes *= 10
				pRes += int(ch - '0')
				break
			}
		}
		fRes += pRes
	}

	return fRes, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	var lookup []string = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine"}

	fRes := 0
	for _, line := range lines {
		pRes := 0
	pre:
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if util.IsDigit(ch) {
				pRes *= 10
				pRes += int(ch - '0')
				break
			}
			for idx, prefix := range lookup {
				if strings.HasPrefix(line[i:], prefix) {
					pRes *= 10
					pRes += int(idx + 1)
					break pre
				}
			}
		}
	post:
		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if util.IsDigit(ch) {
				pRes *= 10
				pRes += int(ch - '0')
				break
			}
			for idx, suffix := range lookup {
				if strings.HasSuffix(line[:i+1], suffix) {
					pRes *= 10
					pRes += int(idx + 1)
					break post
				}
			}
		}
		fRes += pRes
	}

	return fRes, nil
}

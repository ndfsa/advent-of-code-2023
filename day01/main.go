package day01

import (
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func SolvePart1(filePath string, prealloc int) (uint64, error) {
	input, err := util.ReadLines(filePath, prealloc)

	if err != nil {
		return 0, err
	}

	var fRes uint64 = 0
	for _, line := range input {
		var pRes uint64 = 0
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if isDigit(ch) {
				pRes *= 10
				pRes += uint64(ch - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if isDigit(ch) {
				pRes *= 10
				pRes += uint64(ch - '0')
				break
			}
		}
		fRes += pRes
	}

	return fRes, nil
}

func SolvePart2(filePath string, prealloc int) (uint64, error) {
	input, err := util.ReadLines(filePath, prealloc)

	var literal []string = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine"}

	if err != nil {
		return 0, err
	}

	var fRes uint64 = 0
	for _, line := range input {
		var pRes uint64 = 0
	pre:
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if isDigit(ch) {
				pRes *= 10
				pRes += uint64(ch - '0')
				break
			}
			for idx, prefix := range literal {
				if strings.HasPrefix(line[i:], prefix) {
					pRes *= 10
					pRes += uint64(idx + 1)
					break pre
				}
			}
		}
	post:
		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if isDigit(ch) {
				pRes *= 10
				pRes += uint64(ch - '0')
				break
			}
			for idx, suffix := range literal {
				if strings.HasSuffix(line[:i+1], suffix) {
					pRes *= 10
					pRes += uint64(idx + 1)
					break post
				}
			}
		}
		fRes += pRes
	}

	return fRes, nil
}

package day12

import (
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

const (
	STATE_SPACE = 1 << iota
	STATE_SPRING_MID
	STATE_SPRING_END
	STATE_FINAL
	QUESTION_MARK
	PERIOD
	POUND_SIGN
)

func encodeWord(word string) []int {
	res := []int{}
	for _, ch := range word {
		switch ch {
		case '.':
			res = append(res, PERIOD)
		case '?':
			res = append(res, QUESTION_MARK)
		case '#':
			res = append(res, POUND_SIGN)
		}
	}
	return res
}

func countWays(word, enc string) int {
	var automata []int
	for _, spring := range strings.Split(enc, ",") {
		automata = append(automata, STATE_SPACE)
		num, _ := strconv.Atoi(spring)
		for i := 0; i < num-1; i++ {
			automata = append(automata, STATE_SPRING_MID)
		}
		automata = append(automata, STATE_SPRING_END)
	}
	automata = append(automata, STATE_FINAL)

	states := map[int]int{0: 1}
	encWord := encodeWord(word)
	for _, ch := range encWord {
		next := map[int]int{}
		for state, freq := range states {
			switch ch | automata[state] {
			case STATE_SPACE | PERIOD,
				STATE_FINAL | PERIOD,
				STATE_FINAL | QUESTION_MARK:
				next[state] += freq

			case
				STATE_SPACE | QUESTION_MARK:
				next[state] += freq
				next[state+1] += freq

			case STATE_SPACE | POUND_SIGN,
				STATE_SPRING_MID | POUND_SIGN,
				STATE_SPRING_END | QUESTION_MARK,
				STATE_SPRING_END | PERIOD,
				STATE_SPRING_MID | QUESTION_MARK:
				next[state+1] += freq
			}
		}
		states = next
	}
	atLength := len(automata)

	return states[atLength-1] + states[atLength-2]
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	res := 0
	for _, line := range lines {
		word, enc, _ := strings.Cut(line, " ")
		res += countWays(word, enc)
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	res := 0
	for _, line := range lines {
		word, enc, _ := strings.Cut(line, " ")
		res += countWays(word+strings.Repeat("?"+word, 4), enc+strings.Repeat(","+enc, 4))
	}

	return res, nil
}

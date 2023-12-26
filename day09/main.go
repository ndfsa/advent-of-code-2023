package day09

import (
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Sequence struct {
	Data  []int
	Order int
}

func (s Sequence) isArithmeticSequence() bool {
	for i := range s.Data[:len(s.Data)-s.Order] {
		if i == 0 {
			continue
		}

		if s.Data[i] != s.Data[i-1] {
			return false
		}
	}
	return true
}

func (s *Sequence) computeDiff() {
	for i := range s.Data[:len(s.Data)-s.Order-1] {
		s.Data[i] = s.Data[i+1] - s.Data[i]
	}
	s.Order++
}

func (s *Sequence) getNext() int {
	res := 0
	for _, num := range s.Data[len(s.Data)-s.Order-1:] {
		res += num
	}
	return res
}

func parseInput(lines []string) []Sequence {
	sequences := make([]Sequence, 0)
	for _, line := range lines {
		current := Sequence{Data: make([]int, 0)}
		for _, elem := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(elem)
			current.Data = append(current.Data, num)
		}
		current.Order = 0
		sequences = append(sequences, current)
	}
	return sequences
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	sequences := parseInput(lines)
	res := 0
	for _, sequence := range sequences {
		sequence.computeDiff()
		for !sequence.isArithmeticSequence() {
			sequence.computeDiff()
		}

		res += sequence.getNext()
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	sequences := parseInput(lines)

	for idx := range sequences {
		data := sequences[idx].Data
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}

	res := 0
	for _, sequence := range sequences {
		sequence.computeDiff()
		for !sequence.isArithmeticSequence() {
			sequence.computeDiff()
		}

		res += sequence.getNext()
	}

	return res, nil
}

package day01

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolvePart1Example(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./example.txt", 142)
}

func TestSolvePart1(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./input.txt", 54573)
}

func TestSolvePart2Example(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./example2.txt", 281)
}

func TestSolvePart2(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./input.txt", 54591)
}

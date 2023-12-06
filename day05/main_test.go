package day05

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolvePart1Example(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./example.txt", 35)
}

func TestSolvePart1(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./input.txt", 662197086)
}

func TestSolvePart2Example(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./example.txt", 46)
}

func TestSolvePart2(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./input.txt", 0)
}

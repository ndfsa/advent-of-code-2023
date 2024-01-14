package day24

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolvePart1Example(t *testing.T) {
	util.RunSolution(t, SolvePart1(7, 27), "./example.txt", 2)
}

func TestSolvePart1(t *testing.T) {
	util.RunSolution(t, SolvePart1(200000000000000, 400000000000000), "./input.txt", 16589)
}

func TestSolvePart2Example(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./example.txt", 47)
}

func TestSolvePart2(t *testing.T) {
	t.Skip()
	util.RunSolution(t, SolvePart2, "./input.txt", 781390555762385)
}

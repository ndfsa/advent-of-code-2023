package day21

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolvePart1Example(t *testing.T) {
	util.RunSolution(t, SolvePart1(6), "./example.txt", 16)
}

func TestSolvePart1(t *testing.T) {
	util.RunSolution(t, SolvePart1(64), "./input.txt", 3591)
}

func TestSolvePart2Example2(t *testing.T) {
	util.RunSolution(t, SolvePart2(6), "./example.txt", 16)
}

func TestSolvePart2Example3(t *testing.T) {
	util.RunSolution(t, SolvePart2(50), "./example.txt", 1594)
}

func TestSolvePart2Example4(t *testing.T) {
	util.RunSolution(t, SolvePart2(100), "./example.txt", 6536)
}

func TestSolvePart2(t *testing.T) {
	util.RunSolution(t, SolvePart2(65), "./input.txt", 3726)
	util.RunSolution(t, SolvePart2(65+131), "./input.txt", 33086)
	util.RunSolution(t, SolvePart2(65+131+131), "./input.txt", 91672)
	t.Log("use octave solver")
}

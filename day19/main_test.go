package day19

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolvePart1Example(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./example.txt", 19114)
}

func TestSolvePart1(t *testing.T) {
	util.RunSolution(t, SolvePart1, "./input.txt", 398527)
}

func TestSolvePart2Example(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./example.txt", 167409079868000)
}

func TestSolvePart2(t *testing.T) {
	util.RunSolution(t, SolvePart2, "./input.txt", 133973513090020)
}

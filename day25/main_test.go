package day24

import (
	"testing"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func TestSolveExample(t *testing.T) {
	util.RunSolution(t, Solve, "./example.txt", 54)
}

func TestSolve(t *testing.T) {
	util.RunSolution(t, Solve, "./input.txt", 600225)
}

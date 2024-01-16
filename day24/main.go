package day24

import (
	"errors"
	"fmt"
	"math"

	"github.com/ndfsa/advent-of-code-2023/util"
	"gonum.org/v1/gonum/mat"
)

type HailStone struct {
	P util.FVec3
	V util.FVec3
}

func parseInput(lines []string) []HailStone {
	hailStones := []HailStone{}
	for _, line := range lines {
		var px, py, pz, vx, vy, vz int

		fmt.Sscanf(line, "%d,%d,%d @ %d,%d,%d",
			&px,
			&py,
			&pz,
			&vx,
			&vy,
			&vz,
		)

		hailStones = append(hailStones, HailStone{
			P: util.FVec3{X: float64(px), Y: float64(py), Z: float64(pz)},
			V: util.FVec3{X: float64(vx), Y: float64(vy), Z: float64(vz)}})
	}

	return hailStones
}

func SolvePart1(mn, mx float64) util.Solution[int] {
	return util.Solution[int](func(filePath string) (int, error) {
		return solveP1(filePath, mn, mx)
	})
}

func solveP1(filePath string, mn, mx float64) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	hailStones := parseInput(lines)

	res := 0
	for i := 0; i < len(hailStones); i++ {
		for j := i + 1; j < len(hailStones); j++ {
			r := hailStones[i].P
			s := hailStones[j].P
			v := hailStones[i].V
			u := hailStones[j].V

			A := mat.NewDense(2, 2, []float64{v.X, -u.X, v.Y, -u.Y})
			b := mat.NewVecDense(2, []float64{s.X - r.X, s.Y - r.Y})

			var sol mat.VecDense
			if err := sol.SolveVec(A, b); err != nil ||
				sol.At(0, 0) < 0 ||
				sol.At(1, 0) < 0 {
				continue
			}

			in := r.Add(v.Mult(sol.At(0, 0)))

			if mn <= in.X && in.X <= mx &&
				mn <= in.Y && in.Y <= mx {
				res++
			}
		}
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	hailStones := parseInput(lines)

	A := mat.NewDense(10, 10, nil)
	b := mat.NewVecDense(10, nil)

	for i := 0; i < 10; i += 2 {
		p := hailStones[i/2].P
		u := hailStones[i/2].V
		A.SetRow(i, []float64{-u.Y, u.X, 0, p.Y, -p.X, 0, -1, 1, 0, 0})
		b.SetVec(i, u.X*p.Y-u.Y*p.X)

		A.SetRow(i+1, []float64{0, -u.Z, u.Y, 0, p.Z, -p.Y, 0, 0, -1, 1})
		b.SetVec(i+1, u.Y*p.Z-u.Z*p.Y)
	}

	var svd mat.SVD
	if ok := svd.Factorize(A, mat.SVDFull); !ok {
		fmt.Println("failed to factorize")
	}

	if rank := svd.Rank(1e-50); rank != 0 {
		var sol mat.Dense
		svd.SolveTo(&sol, b, rank)

		return int(math.Round(sol.At(0, 0) + sol.At(1, 0) + sol.At(2, 0))), nil

	}

	return 0, errors.New("zero rank system")
}

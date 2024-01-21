package day24

import (
	"fmt"

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

	A := mat.NewDense(4, 4, nil)
	b := mat.NewVecDense(4, nil)

	for i := 0; i < 4; i++ {
		p1 := hailStones[i].P
		u1 := hailStones[i].V
		p2 := hailStones[i+1].P
		u2 := hailStones[i+1].V

		A.SetRow(i, []float64{-u1.Y + u2.Y, u1.X - u2.X, p1.Y - p2.Y, -p1.X + p2.X})
		b.SetVec(i, u1.X*p1.Y-u1.Y*p1.X-u2.X*p2.Y+u2.Y*p2.X)
	}

	var solXY mat.VecDense
	if err := solXY.SolveVec(A, b); err != nil {
		return 0, err
	}

	A = mat.NewDense(2, 2, nil)
	b = mat.NewVecDense(2, nil)
	for i := 0; i < 2; i++ {
		kx := solXY.At(0, 0)
		vx := solXY.At(2, 0)

		px := hailStones[i].P.X
		pz := hailStones[i].P.Z

		ux := hailStones[i].V.X
		uz := hailStones[i].V.Z

		tn := (px - kx) / (vx - ux)

		A.SetRow(i, []float64{1, tn})
		b.SetVec(i, pz+uz*tn)
	}

	var solZ mat.VecDense
	if err := solZ.SolveVec(A, b); err != nil {
		return 0, err
	}

	kx := solXY.At(0, 0)
	ky := solXY.At(1, 0)
	kz := solZ.At(0, 0)

	return int(kx + ky + kz), nil
}

package day24

import (
	"fmt"
	"math"

	"github.com/ndfsa/advent-of-code-2023/util"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize"
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

func buildObjective(p1, p2, p3, v1, v2, v3 util.FVec3) func([]float64) float64 {
	return func(v []float64) float64 {
		// Unpack variables
		k0x, k0y, k0z, vx, vy, vz, t1, t2, t3 := v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8]

		// Define the system of equations
		eq1 := k0x + vx*t1 - v1.X*t1 - p1.X
		eq2 := k0y + vy*t1 - v1.Y*t1 - p1.Y
		eq3 := k0z + vz*t1 - v1.Z*t1 - p1.Z

		eq4 := k0x + vx*t2 - v2.X*t2 - p2.X
		eq5 := k0y + vy*t2 - v2.Y*t2 - p2.Y
		eq6 := k0z + vz*t2 - v2.Z*t2 - p2.Z

		eq7 := k0x + vx*t3 - v3.X*t3 - p3.X
		eq8 := k0y + vy*t3 - v3.Y*t3 - p3.Y
		eq9 := k0z + vz*t3 - v3.Z*t3 - p3.Z

		// Return the sum of squared errors
		return eq1*eq1 + eq2*eq2 + eq3*eq3 + eq4*eq4 + eq5*eq5 + eq6*eq6 + eq7*eq7 + eq8*eq8 + eq9*eq9
	}
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	hailStones := parseInput(lines)

	P1 := hailStones[0]
	P2 := hailStones[1]
	P3 := hailStones[2]

	p1 := P1.P
	v1 := P1.V

	v2 := P2.V
	p2 := P2.P

	p3 := P3.P
	v3 := P3.V

	initialGuess := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}

	problem := optimize.Problem{Func: buildObjective(p1, p2, p3, v1, v2, v3)}

	result, err := optimize.Minimize(problem, initialGuess, nil, nil)
	if err != nil {
		return 0, err
	}

	return int(math.Round(result.X[0] + result.X[1] + result.X[2])), nil
}

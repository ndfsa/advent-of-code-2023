package day22

import (
	"fmt"
	"slices"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Brick struct {
	Blocks []util.Vec3
}

func (b Brick) Supports(other Brick) bool {
	for _, blk1 := range b.Blocks {
		for _, blk2 := range other.Blocks {
			diff := blk2.Sus(blk1)
			if diff == util.DIR_V3_POS_Z {
				return true
			}
		}
	}
	return false
}

type BrickStack struct {
	data    []Brick
	support map[int][]int
}

func (bs *BrickStack) Fall() {
	rest := map[util.Vec3]int{}

	for idx, brick := range bs.data {
		nextPos := make([]util.Vec3, len(brick.Blocks))
		copy(nextPos, brick.Blocks)

		collided := false
		for !collided {
			for i := range nextPos {
				nextPos[i] = nextPos[i].Add(util.DIR_V3_NEG_Z)
			}

			for _, p := range nextPos {
				if brickIndex, ok := rest[p]; ok {
					if !slices.Contains(bs.support[brickIndex], idx) {
						bs.support[brickIndex] = append(bs.support[brickIndex], idx)
					}
					collided = true
				} else if p.Z < 1 {
					collided = true
				}
			}
		}

		for i := range nextPos {
			nextPos[i] = nextPos[i].Add(util.DIR_V3_POS_Z)
			rest[nextPos[i]] = idx 
		}
		bs.data[idx].Blocks = nextPos
	}
}

func (bs BrickStack) InvSupport() map[int][]int {
	res := map[int][]int{}
	for support, supported := range bs.support {
		for _, i := range supported {
			res[i] = append(res[i], support)
		}
	}

	return res
}

func parseInput(lines []string) BrickStack {
	res := []Brick{}
	for _, line := range lines {
		start, end := util.Vec3{}, util.Vec3{}

		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &start.X, &start.Y, &start.Z, &end.X, &end.Y, &end.Z)

		dir := end.Sus(start).Unit()

		blocks := []util.Vec3{}
		current := start
		for current != end {
			blocks = append(blocks, current)
			current = current.Add(dir)
		}
		blocks = append(blocks, current)

		res = append(res, Brick{blocks})
	}

	slices.SortStableFunc(res, func(a, b Brick) int {
		return a.Blocks[0].Z - b.Blocks[0].Z
	})

	return BrickStack{data: res, support: map[int][]int{}}
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	brickStack := parseInput(lines)
	brickStack.Fall()

	supported := brickStack.InvSupport()

	res := 0
	for i := range brickStack.data {
		if top, ok := brickStack.support[i]; ok {
			safe := true
			for _, s := range top {
				if len(supported[s]) < 2 {
					safe = false
					break
				}
			}
			if safe {
				res++
			}
		} else {
			res++
		}
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	brickStack := parseInput(lines)
	brickStack.Fall()

	supported := brickStack.InvSupport()

	foundation := []int{}
	for i := range brickStack.data {
		if top, ok := brickStack.support[i]; ok {
			for _, s := range top {
				if len(supported[s]) < 2 {
					foundation = append(foundation, i)
					break
				}
			}
		}
	}

	res := 0
	for _, fnd := range foundation {
		queue := []int{fnd}
		visited := map[int]struct{}{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			visited[current] = struct{}{}
			for _, top := range brickStack.support[current] {
				supportsLeft := []int{}
				supports := supported[top]
				for _, topSupport := range supports {
					if _, ok := visited[topSupport]; !ok {
						supportsLeft = append(supportsLeft, topSupport)
					}
				}
				if len(supportsLeft) <= 0 {
					res++
					queue = append(queue, top)
				}
			}
		}
	}

	return res, nil
}

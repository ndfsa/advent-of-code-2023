package day16

import (
	"fmt"

	"github.com/ndfsa/advent-of-code-2023/util"
)

const (
	TYPE_EMPTY       byte = '.'
	TYPE_MIRROR      byte = '/'
	TYPE_MIRROR_BACK byte = '\\'
	TYPE_SPLITTER_H  byte = '-'
	TYPE_SPLITTER_V  byte = '|'
)

var (
	DIR_UP    = util.DIR_V2_NEG_X
	DIR_DOWN  = util.DIR_V2_POS_X
	DIR_RIGHT = util.DIR_V2_POS_Y
	DIR_LEFT  = util.DIR_V2_NEG_Y
)

type Photon struct {
	pos util.Vec2
	dir util.Vec2
}

func (p Photon) String() string {
	res := fmt.Sprintf("{%v ", p.pos)
	switch p.dir {
	case DIR_UP:
		res += "U}"
	case DIR_DOWN:
		res += "D}"
	case DIR_LEFT:
		res += "L}"
	case DIR_RIGHT:
		res += "R}"
	}
	return res
}

func (l Photon) move(field [][]byte) []Photon {
	newPhotons := []Photon{}
	instruction := field[l.pos.X][l.pos.Y]

	switch instruction {
	case TYPE_EMPTY:
		l.pos = l.pos.Add(l.dir)
		newPhotons = append(newPhotons, l)
	case TYPE_MIRROR:
		l.dir = util.Vec2{X: -l.dir.Y, Y: -l.dir.X}
		l.pos = l.pos.Add(l.dir)
		newPhotons = append(newPhotons, l)
	case TYPE_MIRROR_BACK:
		l.dir = util.Vec2{X: l.dir.Y, Y: l.dir.X}
		l.pos = l.pos.Add(l.dir)
		newPhotons = append(newPhotons, l)
	case TYPE_SPLITTER_H:
		switch l.dir {
		case DIR_UP, DIR_DOWN:
			newPhotons = append(newPhotons, Photon{pos: l.pos.Add(DIR_LEFT), dir: DIR_LEFT})
			newPhotons = append(newPhotons, Photon{pos: l.pos.Add(DIR_RIGHT), dir: DIR_RIGHT})
		case DIR_RIGHT, DIR_LEFT:
			l.pos = l.pos.Add(l.dir)
			newPhotons = append(newPhotons, l)
		}
	case TYPE_SPLITTER_V:
		switch l.dir {
		case DIR_UP, DIR_DOWN:
			l.pos = l.pos.Add(l.dir)
			newPhotons = append(newPhotons, l)
		case DIR_RIGHT, DIR_LEFT:
			newPhotons = append(newPhotons, Photon{pos: l.pos.Add(DIR_UP), dir: DIR_UP})
			newPhotons = append(newPhotons, Photon{pos: l.pos.Add(DIR_DOWN), dir: DIR_DOWN})
		}
	}

	res := []Photon{}
	for _, photon := range newPhotons {
		if 0 <= photon.pos.X &&
			photon.pos.X < len(field) &&
			0 <= photon.pos.Y &&
			photon.pos.Y < len(field[0]) {

			res = append(res, photon)
		}
	}

	return res
}

func parseInput(lines []string) [][]byte {
	res := [][]byte{}
	for _, line := range lines {
		res = append(res, []byte(line))
	}

	return res
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	field := parseInput(lines)

	queue := map[Photon]struct{}{{pos: util.Vec2{X: 0, Y: 0}, dir: DIR_RIGHT}: {}}
	next := map[Photon][]Photon{}

	for len(queue) > 0 {
		nextQueue := map[Photon]struct{}{}
		for photon := range queue {
			if next[photon] != nil {
				continue
			}
			next[photon] = photon.move(field)
			for _, newPhoton := range next[photon] {
				nextQueue[newPhoton] = struct{}{}
			}
		}

		queue = nextQueue
	}

	energized := make(map[util.Vec2]struct{}, len(next))
	for photon := range next {
		energized[photon.pos] = struct{}{}
	}

	return len(energized), nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	field := parseInput(lines)

	startPhotons := []Photon{}
	for i := 0; i < len(field); i++ {
		startPhotons = append(startPhotons,
			Photon{pos: util.Vec2{X: i, Y: 0}, dir: DIR_RIGHT})
		startPhotons = append(startPhotons,
			Photon{pos: util.Vec2{X: i, Y: len(field[0]) - 1}, dir: DIR_LEFT})
	}
	for i := 0; i < len(field[0]); i++ {
		startPhotons = append(startPhotons,
			Photon{pos: util.Vec2{X: 0, Y: i}, dir: DIR_DOWN})
		startPhotons = append(startPhotons,
			Photon{pos: util.Vec2{X: len(field) - 1, Y: i}, dir: DIR_UP})
	}

	cache := map[Photon][]Photon{}

	res := 0
	for _, start := range startPhotons {
		queue := map[Photon]struct{}{start: {}}

		energized := make(map[Photon]struct{}, 0)
		energized[start] = struct{}{}

		for len(queue) > 0 {
			nextQueue := map[Photon]struct{}{}
			for photon := range queue {
				if cache[photon] == nil {
					cache[photon] = photon.move(field)
				}
				for _, newPhoton := range cache[photon] {
					if _, ok := energized[newPhoton]; ok {
						continue
					}
					nextQueue[newPhoton] = struct{}{}
					energized[newPhoton] = struct{}{}
				}
			}

			queue = nextQueue
		}

		dedup := map[util.Vec2]struct{}{}
		for photon := range energized {
			dedup[photon.pos] = struct{}{}
		}

		if len(dedup) > res {
			res = len(dedup)
		}
	}

	return res, nil
}

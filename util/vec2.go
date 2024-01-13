package util

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Neighbors(accept func(v Vec2, dir Vec2) bool) []Vec2 {
	next := []Vec2{
		DIR_V2_NEG_X,
		DIR_V2_POS_Y,
		DIR_V2_POS_X,
		DIR_V2_NEG_Y,
	}

	res := []Vec2{}
	for _, dir := range next {
		next := v.Add(dir)
		if accept(next, dir) {
			res = append(res, next)
		}
	}
	return res
}

var DIR_V2_NEG_X = Vec2{X: -1, Y: 0}
var DIR_V2_POS_X = Vec2{X: 1, Y: 0}
var DIR_V2_POS_Y = Vec2{X: 0, Y: 1}
var DIR_V2_NEG_Y = Vec2{X: 0, Y: -1}

func (p Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Vec2) AddMult(other Vec2, mag int) Vec2 {
	return Vec2{X: p.X + other.X*mag, Y: p.Y + other.Y*mag}
}

func (p Vec2) HammingDist(other Vec2) int {
	sus := p.Sus(other)
	res := 0

	if sus.X < 0 {
		res += -sus.X
	} else {
		res += sus.X
	}

	if sus.Y < 0 {
		res += -sus.Y
	} else {
		res += sus.Y
	}

	return res
}

func (p Vec2) Sus(other Vec2) Vec2 {
	return Vec2{X: p.X - other.X, Y: p.Y - other.Y}
}

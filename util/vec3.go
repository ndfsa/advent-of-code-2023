package util

import "math"

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) Neighbors(accept func(Vec3) bool) []Vec3 {
	next := []Vec3{
		DIR_V3_POS_X,
		DIR_V3_POS_Y,
		DIR_V3_POS_Z,
		DIR_V3_NEG_X,
		DIR_V3_NEG_Y,
		DIR_V3_NEG_Z,
	}

	res := []Vec3{}
	for _, dir := range next {
		next := v.Add(dir)
		if accept(next) {
			res = append(res, next)
		}
	}
	return res
}

func (v Vec3) Unit() Vec3 {
	if v.X != 0 {
		v.X /= int(math.Abs(float64(v.X)))
	}
	if v.Y != 0 {
		v.Y /= int(math.Abs(float64(v.Y)))
	}
	if v.Z != 0 {
		v.Z /= int(math.Abs(float64(v.Z)))
	}
	return v
}

var DIR_V3_POS_X = Vec3{X: 1, Y: 0, Z: 0}
var DIR_V3_POS_Y = Vec3{X: 0, Y: 1, Z: 0}
var DIR_V3_POS_Z = Vec3{X: 0, Y: 0, Z: 1}
var DIR_V3_NEG_X = Vec3{X: -1, Y: 0, Z: 0}
var DIR_V3_NEG_Y = Vec3{X: 0, Y: -1, Z: 0}
var DIR_V3_NEG_Z = Vec3{X: 0, Y: 0, Z: -1}

func (p Vec3) Add(other Vec3) Vec3 {
	return Vec3{X: p.X + other.X, Y: p.Y + other.Y, Z: p.Z + other.Z}
}

func (p Vec3) AddMult(other Vec3, mag int) Vec3 {
	return Vec3{X: p.X + other.X*mag, Y: p.Y + other.Y*mag, Z: p.Z + other.Z*mag}
}

func (p Vec3) HammingDist(other Vec3) int {
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

	if sus.Z < 0 {
		res += -sus.Z
	} else {
		res += sus.Z
	}

	return res
}

func (p Vec3) Sus(other Vec3) Vec3 {
	return Vec3{X: p.X - other.X, Y: p.Y - other.Y, Z: p.Z - other.Z}
}

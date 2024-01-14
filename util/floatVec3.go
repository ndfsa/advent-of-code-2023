package util

type FVec3 struct {
	X float64
	Y float64
	Z float64
}

func (v FVec3) Mult(scalar float64) FVec3 {
	return FVec3{
		v.X * scalar,
		v.Y * scalar,
		v.Z * scalar}
}

func (p FVec3) Add(other FVec3) FVec3 {
	return FVec3{X: p.X + other.X, Y: p.Y + other.Y, Z: p.Z + other.Z}
}

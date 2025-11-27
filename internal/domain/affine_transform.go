package domain

func AffineTransform(point Point, affineParam AffineParam) Point {
	x := point.X*affineParam.A + point.Y*affineParam.B + affineParam.C
	y := point.X*affineParam.D + point.Y*affineParam.E + affineParam.F
	return NewPoint(x, y)
}

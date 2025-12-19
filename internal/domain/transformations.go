package domain

import "math"

type Transformation string

const (
	Swirl      Transformation = "swirl"
	Horseshoe  Transformation = "horseshoe"
	Sinusoidal Transformation = "sinusoidal"
	Spherical  Transformation = "spherical"
	Heart      Transformation = "heart"
	Cosine     Transformation = "cosine"
)

var AvailableTransformation = map[Transformation]TransFunc{
	Swirl:      swirl,
	Horseshoe:  horseshoe,
	Sinusoidal: sinusoidal,
	Spherical:  spherical,
	Heart:      heart,
	Cosine:     cosine,
}

func (t Transformation) GetTransformation() (TransFunc, bool) {
	fn, ok := AvailableTransformation[t]
	return fn, ok
}

type TransFunc func(point Point) Point

func swirl(point Point) Point {
	x, y := point.X, point.Y

	r2 := x*x + y*y
	s := math.Sin(r2)
	c := math.Cos(r2)
	return NewPoint(x*c-y*s, x*s+y*c)
}

func horseshoe(point Point) Point {
	x, y := point.X, point.Y

	r := math.Hypot(x, y)
	if r == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint((x-y)*(x+y)/r, 2*x*y/r)
}

func spherical(point Point) Point {
	x, y := point.X, point.Y

	r2 := x*x + y*y
	if r2 == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint(x/r2, y/r2)
}

func sinusoidal(point Point) Point {
	x, y := point.X, point.Y

	return NewPoint(math.Sin(x), math.Sin(y))
}

func heart(point Point) Point {
	r := math.Sqrt(point.X*point.X + point.Y*point.Y)
	tetha := math.Atan2(point.X, point.Y)
	newX := r * math.Sin(tetha*r)
	newY := -r * math.Cos(tetha*r)

	return NewPoint(newX, newY)
}

func cosine(point Point) Point {
	newX := math.Cos(math.Pi*point.X) * math.Cosh(point.Y)
	newY := -math.Sin(math.Pi*point.X) * math.Sinh(point.Y)

	return NewPoint(newX, newY)
}

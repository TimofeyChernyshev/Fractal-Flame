package domain

import "math"

type Transformations string

const (
	Swirl      Transformations = "swirl"
	Horseshoe  Transformations = "horseshoe"
	Sinusoidal Transformations = "sinusoidal"
	Spherical  Transformations = "spherical"
)

var AvailableTransformations = map[Transformations]TransFunc{
	Swirl:      swirl,
	Horseshoe:  horseshoe,
	Sinusoidal: sinusoidal,
	Spherical:  spherical,
}

func (t Transformations) GetTransformation() (TransFunc, bool) {
	fn, ok := AvailableTransformations[t]
	return fn, ok
}

type TransFunc func(x, y float64) Point

func swirl(x, y float64) Point {
	r2 := x*x + y*y
	s := math.Sin(r2)
	c := math.Cos(r2)
	return NewPoint(x*c-y*s, x*s+y*c)
}

func horseshoe(x, y float64) Point {
	r := math.Hypot(x, y)
	if r == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint((x-y)*(x+y)/r, 2*x*y/r)
}

func spherical(x, y float64) Point {
	r2 := x*x + y*y
	if r2 == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint(x/r2, y/r2)
}

func sinusoidal(x, y float64) Point {
	return NewPoint(math.Sin(x), math.Sin(y))
}

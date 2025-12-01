package domain

import "math"

// Point представляет точку фрактала
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func Rotate(point Point, theta float64) Point {
	x := point.X
	y := point.Y
	xRotated := x*math.Cos(theta) - y*math.Sin(theta)
	yRotated := x*math.Sin(theta) + y*math.Cos(theta)

	return Point{X: xRotated, Y: yRotated}
}

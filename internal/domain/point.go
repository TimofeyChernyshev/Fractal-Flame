package domain

import (
	"math"
)

// Point представляет точку фрактала
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (point Point) Rotate(theta float64) Point {
	x := point.X
	y := point.Y
	xRotated := x*math.Cos(theta) - y*math.Sin(theta)
	yRotated := x*math.Sin(theta) + y*math.Cos(theta)

	return NewPoint(xRotated, yRotated)
}

// mapPoint мапит точку в пиксель
func (point Point) MapPoint(fi *FractalImage, rect Rectangle) (*Pixel, bool) {
	x := int((point.X - rect.XMin) / rect.Width * float64(fi.Width))
	y := int((point.Y - rect.YMin) / rect.Height * float64(fi.Height))

	return fi.GetPixel(x, y)
}

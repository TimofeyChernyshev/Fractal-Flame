package domain

import "math/rand"

// Rectangle представляет область, в которой находится фрактал
type Rectangle struct {
	X      float64 // минимальное значение X
	Y      float64 // минимальное значение Y
	Width  float64 // ширина квадрата
	Height float64 // высота квадрата
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{X: x, Y: y, Width: width, Height: height}
}

// RandomPoint возвращает случайную точку в границах квадрата
func (r Rectangle) RandomPoint(rnd *rand.Rand) Point {
	x := rnd.Float64()*r.Width + r.X
	y := rnd.Float64()*r.Height + r.Y
	return NewPoint(x, y)
}

// Contains проверяет находится ли точка в границах квадрата
func (r Rectangle) Contains(point Point) bool {
	x, y := point.X, point.Y
	return x >= r.X && x <= r.X+r.Width && y >= r.Y && y <= r.Y+r.Height
}

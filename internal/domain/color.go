package domain

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"

type Color struct {
	R, G, B uint32
}

func NewColor(r, g, b uint32) Color {
	return Color{R: r, G: g, B: b}
}

func RandomColors(rnd random.Random, count int) []Color {
	colors := make([]Color, count)
	for i := range colors {
		colors[i] = NewColor(uint32(rnd.Intn(255)), uint32(rnd.Intn(255)), uint32(rnd.Intn(255)))
	}

	return colors
}

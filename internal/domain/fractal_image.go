package domain

import (
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

const (
	Shift        = 20 // количество итераций для нахождения начальной точки
	IterPerPoint = 50 // количество итераций, при которых прорисовываем точку
)

// FractalImage представляет изображение фрактала
type FractalImage struct {
	Width  int
	Height int
	Pixels []Pixel
}

// NewFractalImage возвращает новый экземпляр FractalImage с указанными шириной и высотой
func NewFractalImage(width, height int) *FractalImage {
	pixels := make([]Pixel, width*height)
	return &FractalImage{Width: width, Height: height, Pixels: pixels}
}

// GetPixel возвращает пиксель, который находится в точке x,y, и флаг того найден он или нет
func (f *FractalImage) GetPixel(x, y int) (*Pixel, bool) {
	if !f.contains(x, y) {
		return nil, false
	}

	return &f.Pixels[y*f.Width+x], true
}

// contains проверяет находится ли пиксель в пределах изображения
func (f *FractalImage) contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.Width && y < f.Height
}

// GenerateFractal создает фрактал
func (f *FractalImage) GenerateFractal(
	rect Rectangle,
	args *Args,
	colors []Color,
	totalFuncWeight float64,
	rnd random.Random,
	iters int,
) {
	for range iters {
		point := rect.RandomPoint(rnd)

		for j := 0; j < Shift+IterPerPoint; j++ {
			affineIndex := rnd.Intn(len(args.AffineParams))
			point = affineTransform(point, args.AffineParams[affineIndex])
			affineParamColor := colors[affineIndex]

			index := getWeightedFunctionIndex(rnd, totalFuncWeight, args.Functions)
			transformation, _ := args.Functions[index].Name.GetTransformation()
			point = transformation(point)

			if j >= Shift {
				theta := 0.0

				for range args.SymmetryLevel {
					rotated := point.Rotate(theta)

					if rect.Contains(rotated) {
						if pixel, ok := rotated.MapPoint(f, rect); ok {
							pixel.ColorPixel(affineParamColor)
						}
					}

					theta += (2 * math.Pi) / float64(args.SymmetryLevel)
				}
			}
		}
	}
}

// affineTransform применять аффинные преобразования на точку
func affineTransform(point Point, affineParam AffineParam) Point {
	x := point.X*affineParam.A + point.Y*affineParam.B + affineParam.C
	y := point.X*affineParam.D + point.Y*affineParam.E + affineParam.F
	return NewPoint(x, y)
}

// getWeightedFunctionIndex возвращает индкекс случайной функции основываясь на общем весе всех функций
func getWeightedFunctionIndex(rnd random.Random, totalWeight float64, functions []Function) int {
	weight := rnd.Float64() * totalWeight
	weightSum := 0.0
	for i, f := range functions {
		weightSum += f.Weight
		if weightSum >= weight {
			return i
		}
	}
	return len(functions) - 1
}

// GammaCorrection применяет гамма-коррекцию на изображение
func (f *FractalImage) GammaCorrection(gamma float64) {
	maxNormal := 0.0

	for y := range f.Height {
		for x := range f.Width {
			pixel, _ := f.GetPixel(x, y)
			if pixel.HitCount > 0 {
				pixel.Normal = math.Log10(float64(pixel.HitCount))
				if pixel.Normal > maxNormal {
					maxNormal = pixel.Normal
				}
			}
		}
	}

	if maxNormal == 0 {
		return
	}

	for y := range f.Height {
		for x := range f.Width {
			pixel, _ := f.GetPixel(x, y)
			if pixel.HitCount > 0 {
				pixel.Normal /= maxNormal
				scale := math.Pow(pixel.Normal, 1.0/gamma)

				pixel.Color.R = uint32(float64(pixel.Color.R) * scale)
				pixel.Color.G = uint32(float64(pixel.Color.G) * scale)
				pixel.Color.B = uint32(float64(pixel.Color.B) * scale)
			}
		}
	}
}

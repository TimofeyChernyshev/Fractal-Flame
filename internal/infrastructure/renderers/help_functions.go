package renderers

import (
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

// getWeightedFunctionIndex возвращает индкекс случайной функции основываясь на общем весе всех функций
func getWeightedFunctionIndex(rnd random.Random, totalWeight float64, functions []domain.Function) int {
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

// mapPoint мапит точку в пиксель
func mapPoint(point domain.Point, fi *domain.FractalImage, rect domain.Rectangle) (*domain.Pixel, bool) {
	x := int((point.X - rect.X) / rect.Width * float64(fi.Width))
	y := int((point.Y - rect.Y) / rect.Height * float64(fi.Height))

	return fi.GetPixel(x, y)
}

func renderIterations(
	rect domain.Rectangle,
	args *domain.Args,
	colors []domain.Color,
	totalFuncWeight float64,
	image *domain.FractalImage,
	rnd random.Random,
	iters int,
) {
	for range iters {
		point := rect.RandomPoint(rnd)

		for j := 0; j < shift+iterPerPoint; j++ {
			affineIndex := rnd.Intn(len(args.AffineParams))
			point = domain.AffineTransform(point, args.AffineParams[affineIndex])
			affineParamColor := colors[affineIndex]

			index := getWeightedFunctionIndex(rnd, totalFuncWeight, args.Functions)
			transformation, _ := args.Functions[index].Name.GetTransformation()
			point = transformation(point)

			if j >= shift {
				theta := 0.0

				for range args.SymmetryLevel {
					rotated := domain.Rotate(point, theta)

					if rect.Contains(rotated) {
						if pixel, ok := mapPoint(rotated, image, rect); ok {
							pixel.ColorPixel(affineParamColor)
						}
					}

					theta += (2 * math.Pi) / float64(args.SymmetryLevel)
				}
			}
		}
	}
}

package renderers

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

// SingleThreadRenderer представляет однопоточную реализацию Renderer
type SingleThreadRenderer struct {
	rect domain.Rectangle
	rnd  random.Random
}

// NewSingleThreadRenderer возвращает новый однопоточный рендерер
func NewSingleThreadRenderer(rnd random.Random) *SingleThreadRenderer {
	return &SingleThreadRenderer{
		rect: domain.NewRectangle(minX, minY, maxX-minX, maxY-minY),
		rnd:  rnd,
	}
}

func (r *SingleThreadRenderer) Render(args *domain.Args) *domain.FractalImage {
	colors := domain.RandomColors(r.rnd, len(args.Functions))

	fractalImage := domain.NewFractalImage(args.Size.Width, args.Size.Height)

	var totalFuncWeight float64
	for _, f := range args.Functions {
		totalFuncWeight += f.Weight
	}

	for range args.IterationCount {
		point := r.rect.RandomPoint(r.rnd)
		for j := range shift + iterForPoint {
			point = domain.AffineTransform(point, args.AffineParams)

			index := getWeightedFunctionIndex(r.rnd, totalFuncWeight, args.Functions)

			functionColor := colors[index]
			transformation, _ := args.Functions[index].Name.GetTransformation()
			point = transformation(point)

			if j >= shift {
				if r.rect.Contains(point) {
					pixel, ok := mapPoint(point, fractalImage, r.rect)
					if ok {
						pixel.ColorPixel(functionColor)
					}
				}
			}
		}
	}
	return fractalImage
}

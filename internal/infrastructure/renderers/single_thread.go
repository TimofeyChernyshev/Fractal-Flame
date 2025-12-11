package renderers

import (
	"log/slog"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// SingleThreadRenderer представляет однопоточную реализацию Renderer
type SingleThreadRenderer struct {
	rect   domain.Rectangle
	rndGen RandomGenerator
}

// NewSingleThreadRenderer возвращает новый однопоточный рендерер
func NewSingleThreadRenderer(rndGen RandomGenerator) *SingleThreadRenderer {
	return &SingleThreadRenderer{
		rect:   domain.NewRectangle(minX, minY, maxX-minX, maxY-minY),
		rndGen: rndGen,
	}
}

func (r *SingleThreadRenderer) Render(args *domain.Args) *domain.FractalImage {
	slog.Info("Starting single thread renderer")

	seed := int64(math.Float64bits(args.Seed))
	rnd := r.rndGen.New(seed)
	colors := domain.RandomColors(rnd, len(args.AffineParams))

	fractalImage := domain.NewFractalImage(args.Size.Width, args.Size.Height)

	var totalFuncWeight float64
	for _, f := range args.Functions {
		totalFuncWeight += f.Weight
	}

	renderIterations(r.rect, args, colors, totalFuncWeight, fractalImage, rnd, args.IterationCount)

	if args.GammaCorrection {
		gammaCorrection(fractalImage, args.Gamma)
	}

	return fractalImage
}

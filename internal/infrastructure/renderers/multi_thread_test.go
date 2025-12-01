package renderers

import (
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

func TestMultiThreadRenderer_Render(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRnd := random.NewMockRandom(ctrl)
	mockRndGen := NewMockRandomGenerator(ctrl)
	mockRndThread1 := random.NewMockRandom(ctrl)
	mockRndThread2 := random.NewMockRandom(ctrl)

	args := &domain.Args{
		Size:           domain.Size{Width: 5, Height: 5},
		IterationCount: 2,
		Threads:        2,
		Seed:           0.0,
		Functions: []domain.Function{
			{Name: domain.Transformations("swirl"), Weight: 1.0},
		},
		AffineParams: []domain.AffineParam{
			{A: 1, B: 0, C: 0, D: 1, E: 0, F: 0},
		},
	}

	baseSeed := int64(math.Float64bits(args.Seed))
	mockRndGen.EXPECT().New(baseSeed).Return(mockRnd)

	// colors := RandomColors(rnd, 1)
	mockRnd.EXPECT().Intn(255).Return(100).Times(1) // R
	mockRnd.EXPECT().Intn(255).Return(150).Times(1) // G
	mockRnd.EXPECT().Intn(255).Return(200).Times(1) // B

	// рандомы для разных потоков
	mockRndGen.EXPECT().New(baseSeed + 1).Return(mockRndThread1)
	mockRndGen.EXPECT().New(baseSeed + 2).Return(mockRndThread2)

	// point := r.rect.RandomPoint(rnd)
	mockRndThread1.EXPECT().Float64().Return(0.5).Times(2)
	mockRndThread2.EXPECT().Float64().Return(0.5).Times(2)

	const calls = shift + iterPerPoint
	// affineIndex := rnd.Intn(len(args.AffineParams))
	mockRndThread1.EXPECT().Intn(len(args.AffineParams)).Return(0).Times(calls)
	mockRndThread2.EXPECT().Intn(len(args.AffineParams)).Return(0).Times(calls)
	// getWeightedFunctionIndex(r.rnd, totalFuncWeight, args.Functions)
	mockRndThread1.EXPECT().Float64().Return(0.0).Times(calls)
	mockRndThread2.EXPECT().Float64().Return(0.0).Times(calls)

	renderer := NewMultiThreadRenderer(mockRndGen)

	img := renderer.Render(args)
	require.NotNil(t, img)

	coloredPixels := 0
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			px, _ := img.GetPixel(x, y)
			if px.Color.R > 0 || px.Color.G > 0 || px.Color.B > 0 {
				coloredPixels++
			}
		}
	}

	require.NotZero(t, coloredPixels)
}

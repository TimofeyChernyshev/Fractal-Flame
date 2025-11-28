package renderers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

func TestGetWeightedFunctionIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rnd := random.NewMockRandom(ctrl)

	functions := []domain.Function{
		{Weight: 1},
		{Weight: 2},
		{Weight: 3},
	}
	totalWeight := 6.0

	tests := []struct {
		name     string
		randVal  float64
		expected int
	}{
		{"select f0", 0.0, 0},
		{"select f0 upper bound", 0.999 / 6.0, 0},
		{"select f1", 1.5 / 6.0, 1},
		{"select f2", 4.5 / 6.0, 2},
		{"select last explicitly", 5.999 / 6.0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd.EXPECT().Float64().Return(tt.randVal)

			idx := getWeightedFunctionIndex(rnd, totalWeight, functions)
			require.Equal(t, tt.expected, idx)
		})
	}
}

func TestMapPoint(t *testing.T) {
	rect := domain.NewRectangle(-1, -1, 2, 2)
	img := domain.NewFractalImage(100, 100)

	tests := []struct {
		name       string
		pointX     float64
		pointY     float64
		pixelX     int
		pixelY     int
		pixelFound bool
	}{
		{
			name:   "center",
			pointX: 0, pointY: 0,
			pixelX: 50, pixelY: 50,
			pixelFound: true,
		},
		{
			name:   "bottom left",
			pointX: -1, pointY: -1,
			pixelX: 0, pixelY: 0,
			pixelFound: true,
		},
		{
			name:   "upper right",
			pointX: 0.99, pointY: 0.99,
			pixelX: 99, pixelY: 99,
			pixelFound: true,
		},
		{
			name:   "out of borders",
			pointX: 5, pointY: 5,
			pixelFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := domain.NewPoint(tt.pointX, tt.pointY)

			pixel, ok := mapPoint(p, img, rect)
			require.Equal(t, tt.pixelFound, ok)

			if ok {
				expectedPixel, _ := img.GetPixel(tt.pixelX, tt.pixelY)
				require.Equal(t, expectedPixel, pixel)
			} else {
				require.Nil(t, pixel)
			}
		})
	}
}

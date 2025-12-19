package domain

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		point    Point
		theta    float64
		expected Point
	}{
		{
			name:     "rotate zero point by zero angle",
			point:    NewPoint(0, 0),
			theta:    0,
			expected: NewPoint(0, 0),
		},
		{
			name:     "rotate (1,0) by 0 degrees",
			point:    NewPoint(1, 0),
			theta:    0,
			expected: NewPoint(1, 0),
		},
		{
			name:     "rotate (1,0) by 90 degrees (pi/2)",
			point:    NewPoint(1, 0),
			theta:    math.Pi / 2,
			expected: NewPoint(0, 1),
		},
		{
			name:     "rotate (1,0) by 180 degrees (pi)",
			point:    NewPoint(1, 0),
			theta:    math.Pi,
			expected: NewPoint(-1, 0),
		},
		{
			name:     "rotate (1,0) by 270 degrees (3pi/2)",
			point:    NewPoint(1, 0),
			theta:    3 * math.Pi / 2,
			expected: NewPoint(0, -1),
		},
		{
			name:     "rotate (1,0) by 360 degrees (2pi)",
			point:    NewPoint(1, 0),
			theta:    2 * math.Pi,
			expected: NewPoint(1, 0),
		},
		{
			name:     "rotate (0,1) by 90 degrees (pi/2)",
			point:    NewPoint(0, 1),
			theta:    math.Pi / 2,
			expected: NewPoint(-1, 0),
		},
		{
			name:     "rotate (1,1) by 45 degrees (pi/4)",
			point:    NewPoint(1, 1),
			theta:    math.Pi / 4,
			expected: NewPoint(0, math.Sqrt(2)),
		},
		{
			name:  "rotate (2,3) by 30 degrees (pi/6)",
			point: NewPoint(2, 3),
			theta: math.Pi / 6,
			expected: NewPoint(
				2*math.Cos(math.Pi/6)-3*math.Sin(math.Pi/6),
				2*math.Sin(math.Pi/6)+3*math.Cos(math.Pi/6),
			),
		},
		{
			name:     "rotate (1,2) by 2pi radians",
			point:    NewPoint(1, 2),
			theta:    2 * math.Pi,
			expected: NewPoint(1, 2),
		},
		{
			name:  "rotate (-1,-1) by pi/3",
			point: NewPoint(-1, -1),
			theta: math.Pi / 3,
			expected: NewPoint(
				-1*math.Cos(math.Pi/3)-(-1)*math.Sin(math.Pi/3),
				-1*math.Sin(math.Pi/3)+(-1)*math.Cos(math.Pi/3),
			),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.point.Rotate(tt.theta)

			require.InDelta(t, tt.expected.X, result.X, 1e-9,
				"X coordinate mismatch: expected %v, got %v", tt.expected.X, result.X)
			require.InDelta(t, tt.expected.Y, result.Y, 1e-9,
				"Y coordinate mismatch: expected %v, got %v", tt.expected.Y, result.Y)
		})
	}
}

func TestMapPoint(t *testing.T) {
	rect := NewRectangle(-1, -1, 2, 2)
	img := NewFractalImage(100, 100)

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
			p := NewPoint(tt.pointX, tt.pointY)

			pixel, ok := p.MapPoint(img, rect)
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

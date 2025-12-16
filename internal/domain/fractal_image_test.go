package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFractalImageGetPixel(t *testing.T) {
	img := NewFractalImage(10, 5)
	tests := []struct {
		name  string
		x, y  int
		found bool
		pixel *Pixel
	}{
		{
			name:  "correct pixel",
			x:     3,
			y:     2,
			found: true,
			pixel: &img.Pixels[2*10+3],
		},
		{
			name:  "x equal width",
			x:     img.Width,
			y:     0,
			found: false,
			pixel: nil,
		},
		{
			name:  "y equal height",
			x:     0,
			y:     img.Height,
			found: false,
			pixel: nil,
		},
		{
			name:  "y and x in the borders",
			x:     img.Width,
			y:     img.Height,
			found: false,
			pixel: nil,
		},
		{
			name:  "x lower zero",
			x:     -1,
			y:     0,
			found: false,
			pixel: nil,
		},
		{
			name:  "y lower zero",
			x:     0,
			y:     -1,
			found: false,
			pixel: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixel, found := img.GetPixel(tt.x, tt.y)
			require.Equal(t, tt.pixel, pixel)
			require.Equal(t, tt.found, found)
		})
	}
}

func TestGammaCorrection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		imageWidth  int
		imageHeight int
		pixels      []struct {
			x, y     int
			color    Color
			hitCount uint32
		}
		gamma       float64
		expectError bool
	}{
		{
			name:        "basic gamma correction",
			imageWidth:  3,
			imageHeight: 3,
			pixels: []struct {
				x, y     int
				color    Color
				hitCount uint32
			}{
				{1, 1, Color{R: 100, G: 100, B: 100}, 100},
				{2, 2, Color{R: 200, G: 200, B: 200}, 1000},
			},
			gamma:       2.2,
			expectError: false,
		},
		{
			name:        "single pixel",
			imageWidth:  1,
			imageHeight: 1,
			pixels: []struct {
				x, y     int
				color    Color
				hitCount uint32
			}{
				{0, 0, Color{R: 255, G: 255, B: 255}, 500},
			},
			gamma:       1.0,
			expectError: false,
		},
		{
			name:        "various hit counts",
			imageWidth:  2,
			imageHeight: 2,
			pixels: []struct {
				x, y     int
				color    Color
				hitCount uint32
			}{
				{0, 0, Color{R: 100, G: 0, B: 0}, 10},   // log10(10) = 1
				{1, 0, Color{R: 0, G: 100, B: 0}, 100},  // log10(100) = 2
				{0, 1, Color{R: 0, G: 0, B: 100}, 1000}, // log10(1000) = 3
				{1, 1, Color{R: 100, G: 100, B: 100}, 10},
			},
			gamma:       2.0,
			expectError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			img := NewFractalImage(tt.imageWidth, tt.imageHeight)

			for _, p := range tt.pixels {
				pixel, ok := img.GetPixel(p.x, p.y)
				require.True(t, ok)
				pixel.Color = p.color
				pixel.HitCount = p.hitCount
			}

			img.GammaCorrection(tt.gamma)

			if tt.gamma == 1.0 {
				for _, p := range tt.pixels {
					pixel, _ := img.GetPixel(p.x, p.y)
					require.Equal(t, p.color.R, pixel.Color.R, "Color should not change with gamma=1.0")
					require.Equal(t, p.color.G, pixel.Color.G, "Color should not change with gamma=1.0")
					require.Equal(t, p.color.B, pixel.Color.B, "Color should not change with gamma=1.0")
				}
			}

			if tt.gamma != 1.0 && len(tt.pixels) > 0 {
				changed := false
				for _, p := range tt.pixels {
					pixel, _ := img.GetPixel(p.x, p.y)
					if p.hitCount > 0 {
						if pixel.Color.R != p.color.R ||
							pixel.Color.G != p.color.G ||
							pixel.Color.B != p.color.B {
							changed = true
							break
						}
					}
				}
				require.True(t, changed, "Colors should be modified by gamma correction")
			}

			for y := 0; y < img.Height; y++ {
				for x := 0; x < img.Width; x++ {
					pixel, _ := img.GetPixel(x, y)
					if pixel.HitCount > 0 {
						require.GreaterOrEqual(t, pixel.Color.R, uint32(0))
						require.GreaterOrEqual(t, pixel.Color.G, uint32(0))
						require.GreaterOrEqual(t, pixel.Color.B, uint32(0))
					}
				}
			}
		})
	}

}

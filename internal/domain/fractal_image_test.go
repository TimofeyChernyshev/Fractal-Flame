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

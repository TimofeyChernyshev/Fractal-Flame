package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRectangleContains(t *testing.T) {
	t.Parallel()

	r := Rectangle{
		XMin:   -1,
		YMin:   -2,
		Width:  1,
		Height: 2,
	}
	tests := []struct {
		name     string
		point    Point
		contains bool
	}{
		{
			name:     "point in rectangle",
			point:    Point{0, 0},
			contains: true,
		},
		{
			name:     "point on the vertex",
			point:    Point{-1, -2},
			contains: true,
		},
		{
			name:     "point out of rectangle",
			point:    Point{2, 0},
			contains: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contains := r.Contains(tt.point)
			require.Equal(t, tt.contains, contains)
		})
	}
}

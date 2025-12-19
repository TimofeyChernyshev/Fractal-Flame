package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRectangleContains(t *testing.T) {
	r := Rectangle{
		XMin:   -1,
		YMin:   -2,
		Width:  1,
		Height: 2,
	}

	require.True(t, r.Contains(Point{0, 0}))
	require.True(t, r.Contains(Point{-1, -2}))

	require.False(t, r.Contains(Point{2, 0}))
	require.False(t, r.Contains(Point{0, 3}))
}

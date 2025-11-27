package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAffineTransform(t *testing.T) {
	p := Point{X: 1, Y: 2}
	params := AffineParam{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
		F: 6,
	}

	result := AffineTransform(p, params)

	require.InDelta(t, result.X, 1*1+2*2+3, 1e-9)
	require.InDelta(t, result.Y, 1*4+2*5+6, 1e-9)
}

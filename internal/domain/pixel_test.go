package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColorPixel(t *testing.T) {
	p := Pixel{}

	// Первое попадание
	p.ColorPixel(NewColor(100, 150, 200))

	require.Equal(t, uint8(100), p.Color.R)
	require.Equal(t, uint8(150), p.Color.G)
	require.Equal(t, uint8(200), p.Color.B)
	require.Equal(t, 1, p.HitCount)

	// Второе попадание
	p.ColorPixel(NewColor(50, 50, 50))

	require.Equal(t, uint8((100+50)/2), p.Color.R)
	require.Equal(t, uint8((150+50)/2), p.Color.G)
	require.Equal(t, uint8((200+50)/2), p.Color.B)
	require.Equal(t, 2, p.HitCount)
}

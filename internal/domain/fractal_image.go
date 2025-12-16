package domain

import (
	"math"
)

// FractalImage представляет изображение фрактала
type FractalImage struct {
	Width  int
	Height int
	Pixels []Pixel
}

// NewFractalImage возвращает новый экземпляр FractalImage с указанными шириной и высотой
func NewFractalImage(width, height int) *FractalImage {
	pixels := make([]Pixel, width*height)
	return &FractalImage{Width: width, Height: height, Pixels: pixels}
}

// GetPixel возвращает пиксель, который находится в точке x,y, и флаг того найден он или нет
func (f *FractalImage) GetPixel(x, y int) (*Pixel, bool) {
	if !f.contains(x, y) {
		return nil, false
	}

	return &f.Pixels[y*f.Width+x], true
}

// contains проверяет находится ли пиксель в пределах изображения
func (f *FractalImage) contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.Width && y < f.Height
}

// GammaCorrection применяет гамма-коррекцию на изображение
func (f *FractalImage) GammaCorrection(gamma float64) {
	maxNormal := 0.0

	for y := range f.Height {
		for x := range f.Width {
			pixel, _ := f.GetPixel(x, y)
			if pixel.HitCount > 0 {
				pixel.Normal = math.Log10(float64(pixel.HitCount))
				if pixel.Normal > maxNormal {
					maxNormal = pixel.Normal
				}
			}
		}
	}

	if maxNormal == 0 {
		return
	}

	for y := range f.Height {
		for x := range f.Width {
			pixel, _ := f.GetPixel(x, y)
			if pixel.HitCount > 0 {
				pixel.Normal /= maxNormal
				scale := math.Pow(pixel.Normal, 1.0/gamma)

				pixel.Color.R = uint32(float64(pixel.Color.R) * scale)
				pixel.Color.G = uint32(float64(pixel.Color.G) * scale)
				pixel.Color.B = uint32(float64(pixel.Color.B) * scale)
			}
		}
	}
}

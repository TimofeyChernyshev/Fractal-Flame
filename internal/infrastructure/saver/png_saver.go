package saver

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type PngSaver struct {
}

func NewPngSaver() *PngSaver {
	return &PngSaver{}
}

func (s *PngSaver) Save(fractalImage *domain.FractalImage, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, fractalImage.Width, fractalImage.Height))

	for y := range fractalImage.Height {
		for x := range fractalImage.Width {
			pixel, _ := fractalImage.GetPixel(x, y)
			pixelColor := color.RGBA{
				R: uint8(pixel.Color.R),
				G: uint8(pixel.Color.G),
				B: uint8(pixel.Color.B),
				A: 255,
			}

			img.Set(x, y, pixelColor)
		}
	}

	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}

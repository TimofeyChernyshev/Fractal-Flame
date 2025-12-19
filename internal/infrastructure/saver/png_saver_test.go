package saver

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type PngSaverSuite struct {
	suite.Suite
	dir   string
	saver *PngSaver
	img   image.Image
}

func TestPngSaverSuite(t *testing.T) {
	suite.Run(t, new(PngSaverSuite))
}

func (s *PngSaverSuite) SetupTest() {
	s.dir = s.T().TempDir()
	s.saver = NewPngSaver()

	width, height := 2, 2
	fractal := domain.NewFractalImage(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := domain.Color{
				R: uint32(x * 100),
				G: uint32(y * 100),
				B: uint32((x + y) * 50),
			}
			pixel, _ := fractal.GetPixel(x, y)
			pixel.ColorPixel(color)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, fractal.Width, fractal.Height))

	for y := range fractal.Height {
		for x := range fractal.Width {
			pixel, _ := fractal.GetPixel(x, y)
			pixelColor := color.RGBA{
				R: uint8(pixel.Color.R),
				G: uint8(pixel.Color.G),
				B: uint8(pixel.Color.B),
				A: 255,
			}

			img.Set(x, y, pixelColor)
		}
	}

	s.img = img
}

func (s *PngSaverSuite) TestSaveSuccessful() {
	tmpPath := filepath.Join(s.dir, "success_test.png")

	err := s.saver.Save(s.img, tmpPath)
	s.Require().NoError(err)

	_, err = os.Stat(tmpPath)
	s.Require().NoError(err)

	file, err := os.Open(tmpPath)
	s.Require().NoError(err)
	defer func() {
		_ = file.Close()
	}()

	_, err = png.Decode(file)
	s.Require().NoError(err)
}

func (s *PngSaverSuite) TestSaveIncorrectPath() {
	wrongPath := filepath.Join(s.dir, "invalid_path/pic.png")

	err := s.saver.Save(s.img, wrongPath)
	s.Require().Error(err)
}

func (s *PngSaverSuite) TestSaveEmptyImage() {
	tmpPath := filepath.Join(s.dir, "empty_test.png")

	img := image.NewRGBA(image.Rect(0, 0, 0, 0))

	err := s.saver.Save(img, tmpPath)
	s.Require().Error(err)
}

package usecase

import (
	"image"
	"image/color"
	"log/slog"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// FlameService представляет основной сервис приложения
type FlameService struct {
	saver    Saver
	renderer Renderer
}

// NewFlameService возвращает новый экземпляр FlameService
func NewFlameService(s Saver, r Renderer) *FlameService {
	return &FlameService{saver: s, renderer: r}
}

// RenderFlame создает и сохраняет фрактальное пламя
func (s *FlameService) RenderFlame(args *domain.Args) error {
	fractal := s.renderer.Render(args)
	slog.Debug("Image rendered")

	image := fractalToImage(fractal)

	err := s.saver.Save(image, args.OutputPath)
	if err != nil {
		slog.Error("Saving image failed", "path", args.OutputPath, "error", err)
		return err
	}
	slog.Debug("Image saved", "path", args.OutputPath)

	slog.Debug("Flame rendered and saved successfully")
	return nil
}

func fractalToImage(fractal *domain.FractalImage) image.Image {
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

	return img
}

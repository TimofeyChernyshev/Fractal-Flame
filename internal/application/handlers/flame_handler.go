package handlers

import (
	"image"
	"image/color"
	"log/slog"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// FlameHandler представляет хендлер рендеринга изображений
type FlameHandler struct {
	saver    Saver
	renderer Renderer
}

// NewFlameService возвращает новый экземпляр FlameService
func NewFlameHandler(s Saver, r Renderer) *FlameHandler {
	return &FlameHandler{saver: s, renderer: r}
}

// RenderFlame создает и сохраняет фрактальное пламя
func (s *FlameHandler) RenderFlame(args *domain.Args) error {
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

package usecase

import (
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
	image := s.renderer.Render(args)
	slog.Debug("Image rendered")

	err := s.saver.Save(image, args.OutputPath)
	if err != nil {
		slog.Error("Saving image failed", "path", args.OutputPath, "error", err)
		return err
	}
	slog.Debug("Image saved", "path", args.OutputPath)

	slog.Debug("Flame rendered and saved successfully")
	return nil
}

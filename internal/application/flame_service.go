package application

import (
	"log/slog"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// FlameService представляет основной сервис приложения
type FlameService struct {
	logger          *slog.Logger
	saver           Saver
	rendererChooser Chooser
	randomGenerator RandomGenerator
}

// NewFlameService возвращает новый экземпляр FlameService
func NewFlameService(s Saver, rc Chooser, rg RandomGenerator, l *slog.Logger) *FlameService {
	return &FlameService{saver: s, rendererChooser: rc, randomGenerator: rg, logger: l}
}

// RenderFlame создает и сохраняет фрактальное пламя
func (s *FlameService) RenderFlame(args *domain.Args) error {
	seed := int64(math.Float64bits(args.Seed))
	randomGen := s.randomGenerator.New(seed)

	renderer := s.rendererChooser.Choose(args.Threads, randomGen)
	s.logger.Debug("Renderer choosed successfully")

	image := renderer.Render(args)
	s.logger.Debug("Image rendered")

	err := s.saver.Save(image, args.OutputPath)
	if err != nil {
		s.logger.Error("Saving image failed", "path", args.OutputPath, "error", err)
		return err
	}

	s.logger.Debug("Flame rendered and saved successfully")
	return nil
}

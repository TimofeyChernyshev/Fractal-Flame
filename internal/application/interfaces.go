package application

import (
	"log/slog"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type Renderer interface {
	Render(args *domain.Args, logger *slog.Logger) *domain.FractalImage
}

type Saver interface {
	Save(fractalImage *domain.FractalImage, path string) error
}

type Chooser interface {
	Choose(threads int) Renderer
}

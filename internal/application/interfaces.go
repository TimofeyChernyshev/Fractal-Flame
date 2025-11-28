package application

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

type Renderer interface {
	Render(args *domain.Args) *domain.FractalImage
}

type Saver interface {
	Save(fractalImage *domain.FractalImage, path string) error
}

type Chooser interface {
	Choose(threads int, random random.Random) Renderer
}

type RandomGenerator interface {
	New(seed int64) random.Random
}

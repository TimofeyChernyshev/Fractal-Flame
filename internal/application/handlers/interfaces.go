package handlers

import (
	"image"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type Renderer interface {
	Render(args *domain.Args) *domain.FractalImage
}

type Saver interface {
	Save(img image.Image, path string) error
}

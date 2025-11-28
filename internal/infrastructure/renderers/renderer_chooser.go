package renderers

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

type Chooser struct {
}

func NewChooser() *Chooser {
	return &Chooser{}
}

func (c *Chooser) Choose(threads int, rnd random.Random) application.Renderer {
	switch threads {
	case 1:
		return NewSingleThreadRenderer(rnd)
	default:
		return NewMultiThreadRenderer(rnd)
	}
}

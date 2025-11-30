package renderers

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

type RandomGenerator interface {
	New(seed int64) random.Random
}

type Chooser struct {
	rndGen RandomGenerator
}

func NewChooser(rndGen RandomGenerator) *Chooser {
	return &Chooser{
		rndGen: rndGen,
	}
}

func (c *Chooser) Choose(threads int) application.Renderer {
	switch threads {
	case 1:
		return NewSingleThreadRenderer(c.rndGen)
	default:
		return NewMultiThreadRenderer(c.rndGen)
	}
}

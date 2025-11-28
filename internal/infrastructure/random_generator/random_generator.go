package random_generator

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

type RandomGenerator struct {
}

func NewGenerator() *RandomGenerator {
	return &RandomGenerator{}
}

type Randomizer struct {
	r *rand.Rand
}

func (r *RandomGenerator) New(seed int64) random.Random {
	return &Randomizer{
		r: rand.New(rand.NewSource(seed)),
	}
}

func (s *Randomizer) Float64() float64 {
	return s.r.Float64()
}

func (s *Randomizer) Intn(n int) int {
	return s.r.Intn(n)
}

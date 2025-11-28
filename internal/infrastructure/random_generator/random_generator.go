package random_generator

import "math/rand"

type Random struct {
	r *rand.Rand
}

func New(seed int64) *Random {
	return &Random{
		r: rand.New(rand.NewSource(seed)),
	}
}

func (s *Random) Float64() float64 {
	return s.r.Float64()
}

func (s *Random) Intn(n int) int {
	return s.r.Intn(n)
}

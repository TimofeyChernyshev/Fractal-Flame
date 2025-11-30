package renderers

import (
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/random_generator"
)

func createArgs(iterations, threads int) *domain.Args {
	return &domain.Args{
		Size: domain.Size{
			Width:  800,
			Height: 600,
		},
		IterationCount: iterations,
		Threads:        threads,
		Functions: []domain.Function{
			{Name: domain.Sinusoidal, Weight: 1.0},
			{Name: domain.Spherical, Weight: 1.0},
		},
		AffineParams: domain.AffineParam{
			A: 0.5, B: 0.5, C: 0.5,
			D: 0.5, E: 0.5, F: 0.5,
		},
	}
}

func BenchmarkRenderers(b *testing.B) {
	testCases := []struct {
		name       string
		iterations int
		threads    int
	}{
		{"1k", 1_000, 4},
		{"10k", 10_000, 4},
		{"100k", 100_000, 4},
	}

	for _, tc := range testCases {
		b.Run(tc.name+"_SingleThread", func(b *testing.B) {
			args := createArgs(tc.iterations, 1)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				rnd := random_generator.NewGenerator()
				renderer := NewSingleThreadRenderer(rnd)
				renderer.Render(args)
			}
		})

		b.Run(tc.name+"_MultiThread", func(b *testing.B) {
			args := createArgs(tc.iterations, tc.threads)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				rnd := random_generator.NewGenerator()
				renderer := NewMultiThreadRenderer(rnd)
				renderer.Render(args)
			}
		})
	}
}

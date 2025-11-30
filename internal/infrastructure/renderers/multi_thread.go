package renderers

import (
	"math"
	"sync"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// MultiThreadRenderer представляет многопоточную реализацию Renderer
type MultiThreadRenderer struct {
	rect   domain.Rectangle
	rndGen RandomGenerator
}

// NewMultiThreadRenderer возвращает новый многопоточный рендерер
func NewMultiThreadRenderer(rnd RandomGenerator) *MultiThreadRenderer {
	return &MultiThreadRenderer{
		rect:   domain.NewRectangle(minX, minY, maxX-minX, maxY-minY),
		rndGen: rnd,
	}
}

func (r *MultiThreadRenderer) Render(args *domain.Args) *domain.FractalImage {
	baseSeed := int64(math.Float64bits(args.Seed))
	rnd := r.rndGen.New(baseSeed)
	colors := domain.RandomColors(rnd, len(args.Functions))

	var totalFuncWeight float64
	for _, f := range args.Functions {
		totalFuncWeight += f.Weight
	}

	threads := args.Threads
	iterationsPerThread := args.IterationCount / threads

	results := make(chan *domain.FractalImage, threads)
	var wg sync.WaitGroup

	for i := range threads {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			// Локалный генератор для каждого потока
			seed := baseSeed + int64(workerID)
			localRnd := r.rndGen.New(seed)

			// Локальная картинка, чтобы избежать гонки записи
			workerImage := domain.NewFractalImage(args.Size.Width, args.Size.Height)

			startIter := workerID * iterationsPerThread
			endIter := startIter + iterationsPerThread
			if workerID == threads-1 {
				endIter = args.IterationCount
			}

			renderIterations(r.rect, args, colors, totalFuncWeight, workerImage, localRnd, startIter, endIter)

			results <- workerImage
		}(i + 1)
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Объединение результатов
	finalImage := domain.NewFractalImage(args.Size.Width, args.Size.Height)
	for workerImage := range results {
		r.mergeImages(finalImage, workerImage)
	}

	return finalImage
}

// mergeImages объединяет частичные результаты
func (r *MultiThreadRenderer) mergeImages(final, partial *domain.FractalImage) {
	for y := 0; y < final.Height; y++ {
		for x := 0; x < final.Width; x++ {
			finalPixel, _ := final.GetPixel(x, y)
			partialPixel, _ := partial.GetPixel(x, y)

			if partialPixel.HitCount == 0 {
				continue
			}

			if finalPixel.HitCount == 0 {
				finalPixel.Color = partialPixel.Color
				finalPixel.HitCount = partialPixel.HitCount
			} else {
				totalHits := finalPixel.HitCount + partialPixel.HitCount

				finalPixel.ColorPixel(partialPixel.Color)

				finalPixel.HitCount = totalHits
			}
		}
	}
}

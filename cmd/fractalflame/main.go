package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	renderer "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/flame_renderer"
	usecase "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application/flame_usecase"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/cli"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/random_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/saver"
)

func main() {
	randomGen := random_generator.NewGenerator()

	logger := NewLogger()
	slog.SetDefault(logger)

	saver := saver.NewPngSaver()
	renderer := renderer.NewRenderer(randomGen)

	flameService := usecase.NewFlameService(saver, renderer)

	app := cli.NewApp(flameService)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	err := app.Run(ctx, os.Args)
	if err != nil {
		log.Println(err)
	}
}

func NewLogger() *slog.Logger {
	env := os.Getenv("APP_ENV")

	switch env {
	case "prod":
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	default:
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	}
}

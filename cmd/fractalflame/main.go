package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/cli"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/random_generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/renderers"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/infrastructure/saver"
)

func main() {
	randomGen := random_generator.NewGenerator()

	logger := NewLogger()
	slog.SetDefault(logger)

	saver := saver.NewPngSaver()
	chooser := renderers.NewChooser(randomGen)

	flameService := application.NewFlameService(saver, chooser)

	app := cli.NewApp(flameService)

	err := app.Run(context.Background(), os.Args)
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

package cli

import (
	"context"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type FlameService interface {
	RenderFlame(args *domain.Args) error
}

type App struct {
	FlameService FlameService
}

func NewApp(s FlameService) *App {
	return &App{FlameService: s}
}

// Run парсит аргументы командной строки и запускает само приложение
func (a *App) Run(ctx context.Context, args []string) error {
	app := &cli.Command{
		Name:     "fractal-flame",
		Usage:    "Generates fractal flames",
		HideHelp: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "help",
				Usage: "Show help",
			},
			&cli.IntFlag{
				Name:      "width",
				Aliases:   []string{"w"},
				Value:     1920,
				Usage:     "Width of the final image",
				Validator: validateDimension,
			},
			&cli.IntFlag{
				Name:      "height",
				Aliases:   []string{"h"},
				Value:     1080,
				Usage:     "Height of the final image",
				Validator: validateDimension,
			},
			&cli.Float64Flag{
				Name:  "seed",
				Value: 5.1234,
				Usage: "Initial value for the random generator",
			},
			&cli.IntFlag{
				Name:      "iteration-count",
				Aliases:   []string{"i"},
				Value:     2500,
				Usage:     "Number of generation iterations",
				Validator: validateIterationCount,
			},
			&cli.StringFlag{
				Name:      "output-path",
				Aliases:   []string{"o"},
				Value:     "result.png",
				Usage:     "Relative path to the PNG output file",
				Validator: validateOutput,
			},
			&cli.IntFlag{
				Name:      "threads",
				Aliases:   []string{"t"},
				Value:     1,
				Usage:     "Number of threads to use",
				Validator: validateThreads,
			},
			&cli.Float64SliceFlag{
				Name:      "affine-params",
				Aliases:   []string{"ap"},
				Value:     []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
				Usage:     "Affine transform params <a>,<b>,<c>,<d>,<e>,<f>",
				Validator: validateAffineParams,
			},
			&cli.StringSliceFlag{
				Name:      "functions",
				Aliases:   []string{"f"},
				Value:     []string{"swirl:1.0"},
				Usage:     "Transform functions: <func>:<weight>,...",
				Validator: validateFunctions,
			},
			&cli.StringFlag{
				Name:      "config",
				Usage:     "Relative path to json config file",
				Validator: validateConfig,
			},
		},
		Action: a.runApp,
	}

	err := app.Run(ctx, args)
	return err
}

// runApp запускает основной сервис приложения
func (a *App) runApp(_ context.Context, c *cli.Command) error {
	args, err := parseArgs(c)
	if err != nil {
		return err
	}

	err = a.FlameService.RenderFlame(args)
	if err != nil {
		return err
	}
	return nil
}

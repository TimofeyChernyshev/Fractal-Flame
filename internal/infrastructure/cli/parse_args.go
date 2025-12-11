package cli

import (
	"log/slog"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// parseArgs парсит аргументы командной строки
// в порядке консольный ввод, файл конфигурации, дефолтные значения
func (a *App) parseArgs(c *cli.Command) (*domain.Args, error) {
	affine := c.Float64Slice("affine-params")
	// 6 - количество аффинных параметров в одной группе параметров
	affineParamsCount := len(affine) / 6

	args := &domain.Args{
		Size: domain.Size{
			Width:  c.Int("width"),
			Height: c.Int("height"),
		},
		Seed:            c.Float64("seed"),
		IterationCount:  c.Int("iteration-count"),
		OutputPath:      c.String("output-path"),
		Threads:         c.Int("threads"),
		Functions:       parseFunctions(c.StringSlice("functions")),
		GammaCorrection: c.Bool("gamma-correction"),
		Gamma:           c.Float64("gamma"),
		SymmetryLevel:   c.Int("symmetry-level"),
	}

	args.AffineParams = make([]domain.AffineParam, affineParamsCount)
	for i := range affineParamsCount {
		A, B, C, D, E, F := affine[i*6], affine[i*6+1], affine[i*6+2], affine[i*6+3], affine[i*6+4], affine[i*6+5]
		args.AffineParams[i] = domain.AffineParam{A: A, B: B, C: C, D: D, E: E, F: F}
	}

	if c.IsSet("gamma") && !c.IsSet("gamma-correction") {
		args.GammaCorrection = true
	}

	if c.IsSet("config") {
		configPath := c.String("config")

		slog.Debug("Reading config", "path", configPath)

		err := a.readConfig(configPath, c, args)
		if err != nil {
			slog.Error("Failed to read config", "error", err)
			return nil, err
		}
	}

	return args, nil
}

func parseFunctions(funcs []string) []domain.Function {
	var functions []domain.Function
	for _, f := range funcs {
		if function, err := parseFunc(f); err == nil {
			functions = append(functions, function)
		}
	}
	return functions
}

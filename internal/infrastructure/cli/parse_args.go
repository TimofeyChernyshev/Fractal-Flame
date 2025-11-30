package cli

import (
	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

// parseArgs парсит аргументы командной строки
// в порядке консольный ввод, файл конфигурации, дефолтные значения
func (a *App) parseArgs(c *cli.Command) (*domain.Args, error) {
	affine := c.Float64Slice("affine-params")

	args := &domain.Args{
		Size: domain.Size{
			Width:  c.Int("width"),
			Height: c.Int("height"),
		},
		Seed:            c.Float64("seed"),
		IterationCount:  c.Int("iteration-count"),
		OutputPath:      c.String("output-path"),
		Threads:         c.Int("threads"),
		AffineParams:    domain.AffineParam{A: affine[0], B: affine[1], C: affine[2], D: affine[3], E: affine[4], F: affine[5]},
		Functions:       parseFunctions(c.StringSlice("functions")),
		GammaCorrection: c.Bool("gamma-correction"),
		Gamma:           c.Float64("gamma"),
	}

	if c.IsSet("gamma") && !c.IsSet("gamma-correction") {
		args.GammaCorrection = true
	}

	if c.IsSet("config") {
		configPath := c.String("config")

		a.logger.Debug("Reading config", "path", configPath)

		err := a.readConfig(configPath, c, args)
		if err != nil {
			a.logger.Error("Failed to read config", "error", err)
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

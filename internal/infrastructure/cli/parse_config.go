package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

var errConfigArgs error = errors.New("invalid value in config file")

// readConfig считвает конфиг файл и ставит параметрам значения из него
func (a *App) readConfig(configPath string, c *cli.Command, args *domain.Args) error {
	var cfg domain.Args
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		slog.Error("Failed to parse config file", "error", err)
		return fmt.Errorf("cannot parse config file: %w", err)
	}

	fieldMap := []struct {
		cliFlag   string
		configVal interface{}
		validator func(interface{}) error
	}{
		{"width", cfg.Size.Width, func(v interface{}) error { return validateDimension(v.(int)) }},
		{"height", cfg.Size.Height, func(v interface{}) error { return validateDimension(v.(int)) }},
		{"seed", cfg.Seed, nil},
		{"iteration-count", cfg.IterationCount, func(v interface{}) error { return validateIterationCount(v.(int)) }},
		{"output-path", cfg.OutputPath, func(v interface{}) error { return validateOutput(v.(string)) }},
		{"threads", cfg.Threads, func(v interface{}) error { return validateThreads(v.(int)) }},
		{"gamma-correction", cfg.GammaCorrection, nil},
		{"gamma", cfg.Gamma, func(v interface{}) error { return validateGamma(v.(float64)) }},
		{"symmetry-level", cfg.SymmetryLevel, func(v interface{}) error { return validateSymmetryLevel(v.(int)) }},
	}

	for _, field := range fieldMap {
		if !c.IsSet(field.cliFlag) && !isZero(field.configVal) {
			if field.validator != nil {
				if err := field.validator(field.configVal); err != nil {
					slog.Error("Config contains wrong argument value", "flag", field.cliFlag, "value", field.configVal, "error", err)
					return fmt.Errorf("%w: %w", errConfigArgs, err)
				}
			}
			setFieldValue(args, field.cliFlag, field.configVal)
			slog.Debug("Get arg from config", "flag", field.cliFlag, "value", field.configVal)
		}
	}

	if !c.IsSet("functions") && len(cfg.Functions) > 0 {
		for _, f := range cfg.Functions {
			_, ok := domain.Transformation(f.Name).GetTransformation()
			if !ok {
				slog.Error("Provided function isn't supported", "function", f.Name)
				return fmt.Errorf("%w %s: %w: transformation function isn't supported", errConfigArgs, f.Name, errFunctionFormat)
			}
			if f.Weight <= 0 {
				slog.Error("Function weight lower or equal zero", "weight", f.Weight)
				return fmt.Errorf("%w %f: weight must be positive number", errConfigArgs, f.Weight)
			}
		}
		args.Functions = cfg.Functions
	}
	if !c.IsSet("affine-params") && len(cfg.AffineParams) > 0 {
		var affineParams []domain.AffineParam
		for _, p := range cfg.AffineParams {
			if !isZero(p) {
				affineParams = append(affineParams, p)
			}
		}
		args.AffineParams = affineParams
	}

	return nil
}

// setFieldValue присваивает значениям args новые значения основываясь на имени флага
func setFieldValue(args *domain.Args, field string, value interface{}) {
	switch field {
	case "width":
		args.Size.Width = value.(int)
	case "height":
		args.Size.Height = value.(int)
	case "seed":
		args.Seed = value.(float64)
	case "iteration-count":
		args.IterationCount = value.(int)
	case "output-path":
		args.OutputPath = value.(string)
	case "threads":
		args.Threads = value.(int)
	case "gamma-correction":
		args.GammaCorrection = value.(bool)
	case "gamma":
		args.Gamma = value.(float64)
	case "symmetry-level":
		args.SymmetryLevel = value.(int)
	}
}

// isZero проверяет является ли значение параметра нулевым
func isZero(v interface{}) bool {
	switch val := v.(type) {
	case int:
		return val == 0
	case float64:
		return val == 0
	case string:
		return val == ""
	case domain.AffineParam:
		return val.A == 0 && val.B == 0 && val.C == 0 && val.D == 0 && val.E == 0 && val.F == 0
	case bool:
		return v == false
	default:
		return true
	}
}

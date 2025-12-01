package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

var (
	errWrongDimension error = errors.New("dimension is lower or equal zero")
	errIterationCount error = errors.New("iteration count is lower or equal zero")
	errWrongOutput    error = errors.New("cannot create file with output param")
	errThreadsCount   error = errors.New("threads count is lower or equal zero")
	errAffineParams   error = errors.New("wrong amount of affine params")
	errFunctionFormat error = errors.New("function provided in wrong format")
	errConfigExt      error = errors.New("wrong config file extension")
	errWrongGamma     error = errors.New("wrong value for gamma")
)

func validateDimension(d int) error {
	if d < 1 {
		return errWrongDimension
	}
	return nil
}

func validateIterationCount(i int) error {
	if i < 1 {
		return errIterationCount
	}
	return nil
}

func validateOutput(o string) error {
	ext := filepath.Ext(o)
	if ext != ".png" {
		return fmt.Errorf("%w: extension doesn't supported: %s", errWrongOutput, ext)
	}

	if filepath.IsAbs(o) {
		return fmt.Errorf("%w: path isn't relative", errWrongOutput)
	}

	dir := filepath.Dir(o)
	tmpFile, err := os.CreateTemp(dir, ".tmp")
	if err != nil {
		return fmt.Errorf("%w: cannot write file to directory '%s': %w", errWrongOutput, dir, err)
	}
	_ = tmpFile.Close()
	_ = os.Remove(tmpFile.Name())

	return nil
}

func validateThreads(t int) error {
	if t < 1 {
		return errThreadsCount
	}
	return nil
}

func validateAffineParams(ap []float64) error {
	if len(ap)%6 != 0 {
		return fmt.Errorf("%w: %d", errAffineParams, len(ap))
	}
	return nil
}

func validateFunctions(funcs []string) error {
	for _, f := range funcs {
		_, err := parseFunc(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseFunc(f string) (domain.Function, error) {
	functionStr, weightStr, ok := strings.Cut(f, ":")
	if !ok {
		return domain.Function{}, fmt.Errorf("%w: %s", errFunctionFormat, f)
	}

	_, ok = domain.Transformations(functionStr).GetTransformation()
	if !ok {
		return domain.Function{}, fmt.Errorf("%w: transformation function isn't supported: %s", errFunctionFormat, functionStr)
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return domain.Function{}, fmt.Errorf("%w: weight must be a float number: %s", errFunctionFormat, weightStr)
	}
	if weight <= 0 {
		return domain.Function{}, fmt.Errorf("%w: weight must be positive number: %s", errFunctionFormat, weightStr)
	}

	return domain.Function{Name: domain.Transformations(functionStr), Weight: weight}, nil
}

func validateConfig(c string) error {
	if filepath.Ext(c) != ".json" {
		return errConfigExt
	}
	_, err := os.ReadFile(c)
	if err != nil {
		return fmt.Errorf("cannot read config file: %w", err)
	}

	return nil
}

func validateGamma(g float64) error {
	if g == 0 {
		return errWrongGamma
	}

	return nil
}

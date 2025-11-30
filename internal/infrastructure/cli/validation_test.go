package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type validationSuite struct {
	suite.Suite
	config *os.File
}

func TestValidationSuite(t *testing.T) {
	suite.Run(t, new(validationSuite))
}

func (s *validationSuite) SetupSuite() {
	tempDir := s.T().TempDir()

	var err error

	s.config, err = os.CreateTemp(tempDir, "json-*.json")
	s.Require().NoError(err)
	_ = s.config.Close()
}

func (s *validationSuite) TestValidateDimension() {
	tests := []struct {
		name        string
		dimension   int
		expectedErr error
	}{
		{
			name:        "valid dimension",
			dimension:   2,
			expectedErr: nil,
		},
		{
			name:        "dimension lower 0",
			dimension:   -1,
			expectedErr: errWrongDimension,
		},
		{
			name:        "dimension equal 0",
			dimension:   0,
			expectedErr: errWrongDimension,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateDimension(tt.dimension)

			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}

func (s *validationSuite) TestValidateIterationCount() {
	tests := []struct {
		name        string
		itCount     int
		expectedErr error
	}{
		{
			name:        "valid iteration count",
			itCount:     2,
			expectedErr: nil,
		},
		{
			name:        "iteration count lower 0",
			itCount:     -1,
			expectedErr: errIterationCount,
		},
		{
			name:        "iteration count equal 0",
			itCount:     0,
			expectedErr: errIterationCount,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateIterationCount(tt.itCount)

			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}

func (s *validationSuite) TestValidateThreads() {
	tests := []struct {
		name        string
		threads     int
		expectedErr error
	}{
		{
			name:        "valid threads count",
			threads:     2,
			expectedErr: nil,
		},
		{
			name:        "threads count lower 0",
			threads:     -1,
			expectedErr: errThreadsCount,
		},
		{
			name:        "threads count equal 0",
			threads:     0,
			expectedErr: errThreadsCount,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateThreads(tt.threads)

			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}

func (s *validationSuite) TestValidateOutput() {
	absPath, _ := filepath.Abs("y:/out.png")
	tests := []struct {
		name        string
		output      string
		expectedErr string
	}{
		{
			name:        "valid output",
			output:      ".png",
			expectedErr: "",
		},
		{
			name:        "absolute path to output",
			output:      absPath,
			expectedErr: "path isn't relative",
		},
		{
			name:        "wrong extension",
			output:      ".txt",
			expectedErr: "extension doesn't supported",
		},
		{
			name:        "cannot write file to directory",
			output:      "unexistantDirecory/.png",
			expectedErr: "cannot write file to directory",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateOutput(tt.output)

			if err != nil || tt.expectedErr != "" {
				s.Require().ErrorContains(err, tt.expectedErr)
			}
		})
	}
}

func (s *validationSuite) TestValidateAffineParams() {
	tests := []struct {
		name         string
		affineParams []float64
		expectedErr  error
	}{
		{
			name:         "valid affine params",
			affineParams: []float64{0.1, 2, 0.1, -12, 0.4, 2},
			expectedErr:  nil,
		},
		{
			name:         "wrong amount of affine params",
			affineParams: []float64{0.1},
			expectedErr:  errAffineParams,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateAffineParams(tt.affineParams)

			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}

func (s *validationSuite) TestValidateFunctions() {
	tests := []struct {
		name        string
		functions   []string
		expectedErr string
	}{
		{
			name:        "valid funcs",
			functions:   []string{"swirl:1.0", "horseshoe:0.5"},
			expectedErr: "",
		},
		{
			name:        "unvailable func",
			functions:   []string{"func:1.0"},
			expectedErr: "transformation function isn't supported",
		},
		{
			name:        "wrong format",
			functions:   []string{"swirl:1.0", "swirl"},
			expectedErr: errFunctionFormat.Error(),
		},
		{
			name:        "weight isn't float",
			functions:   []string{"swirl:a"},
			expectedErr: "weight must be a float number",
		},
		{
			name:        "weight lower zero",
			functions:   []string{"swirl:-0.1"},
			expectedErr: "weight must be positive number",
		},
		{
			name:        "weight equal zero",
			functions:   []string{"swirl:0"},
			expectedErr: "weight must be positive number",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateFunctions(tt.functions)

			if err != nil || tt.expectedErr != "" {
				s.Require().ErrorContains(err, tt.expectedErr)
			}
		})
	}
}

func (s *validationSuite) TestValidateConfig() {
	tests := []struct {
		name        string
		config      string
		expectedErr string
	}{
		{
			name:        "valid config file",
			config:      s.config.Name(),
			expectedErr: "",
		},
		{
			name:        "wrong config file extension",
			config:      ".txt",
			expectedErr: errConfigExt.Error(),
		},
		{
			name:        "cannot read config file",
			config:      "dir/unexistant_config.json",
			expectedErr: "cannot read config file",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateConfig(tt.config)

			if err != nil || tt.expectedErr != "" {
				s.Require().ErrorContains(err, tt.expectedErr)
			}
		})
	}
}

func (s *validationSuite) TestValidateGamma() {
	tests := []struct {
		name        string
		gamma       float64
		expectedErr error
	}{
		{
			name:        "valid gamma",
			gamma:       2.0,
			expectedErr: nil,
		},
		{
			name:        "wrong gamma",
			gamma:       0,
			expectedErr: errWrongGamma,
		},
		{
			name:        "negative gamma",
			gamma:       -2,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateGamma(tt.gamma)

			s.Require().ErrorIs(err, tt.expectedErr)
		})
	}
}

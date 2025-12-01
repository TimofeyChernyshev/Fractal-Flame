package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type appSuite struct {
	suite.Suite
	service    *MockFlameService
	ctrl       *gomock.Controller
	app        *App
	configFile *os.File
	wrongFile  *os.File
	logBuffer  io.Writer
}

func TestRunAppSuite(t *testing.T) {
	suite.Run(t, new(appSuite))
}

func (s *appSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.service = NewMockFlameService(s.ctrl)

	s.logBuffer = &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(s.logBuffer, &slog.HandlerOptions{Level: slog.LevelError}))

	s.app = NewApp(s.service, logger)

	tempDir := s.T().TempDir()
	var err error
	cfg := domain.Args{
		Size:           domain.Size{Height: 999, Width: 999},
		IterationCount: 999,
		OutputPath:     "cfg.png",
		Threads:        9,
		Seed:           9.99,
		Functions: []domain.Function{
			{Name: domain.Swirl, Weight: 0.5},
			{Name: domain.Horseshoe, Weight: 0.2},
		},
		AffineParams:    []domain.AffineParam{{A: 2, B: 2, C: 2, D: 2, E: 2, F: 2}},
		GammaCorrection: true,
		Gamma:           2,
		SymmetryLevel:   5,
	}
	jsonBytes, err := json.Marshal(cfg)
	s.Require().NoError(err)
	s.configFile, err = os.CreateTemp(tempDir, "config-*.json")
	s.Require().NoError(err)
	s.wrongFile, err = os.CreateTemp(tempDir, "wrong-*.json")
	s.Require().NoError(err)
	_, err = s.configFile.Write(jsonBytes)
	s.Require().NoError(err)

	_ = s.configFile.Close()
	_ = s.wrongFile.Close()
}

func (s *appSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *appSuite) TestParseArgs() {
	tests := []struct {
		name         string
		args         []string
		expectedArgs *domain.Args
	}{
		{
			name: "default values",
			args: []string{"fractal-flame"},
			expectedArgs: &domain.Args{
				Size:            domain.Size{Height: 1080, Width: 1920},
				IterationCount:  2500,
				OutputPath:      "result.png",
				Threads:         1,
				Seed:            5.1234,
				Functions:       []domain.Function{{Name: domain.Swirl, Weight: 1.0}},
				AffineParams:    []domain.AffineParam{{A: 0.9, B: 0.7, C: 0, D: -0.15, E: -1.1, F: 0}},
				GammaCorrection: false,
				Gamma:           2.2,
				SymmetryLevel:   1,
			},
		},
		{
			name: "set all flags",
			args: []string{"fractal-flame", "--height", "1", "--width", "1", "--iteration-count", "1",
				"--output-path", ".png", "--threads", "5", "--seed", "1", "--affine-params", "1,1,1,1,1,1", "--functions", "swirl:0.5,horseshoe:0.1",
				"--gamma-correction", "--gamma", "2", "--symmetry-level", "5",
			},
			expectedArgs: &domain.Args{
				Size:            domain.Size{Height: 1, Width: 1},
				IterationCount:  1,
				OutputPath:      ".png",
				Threads:         5,
				Seed:            1,
				Functions:       []domain.Function{{Name: domain.Swirl, Weight: 0.5}, {Name: domain.Horseshoe, Weight: 0.1}},
				AffineParams:    []domain.AffineParam{{A: 1, B: 1, C: 1, D: 1, E: 1, F: 1}},
				GammaCorrection: true,
				Gamma:           2,
				SymmetryLevel:   5,
			},
		},
		{
			name: "set all flags and config",
			args: []string{"fractal-flame", "--height", "1", "--width", "1", "--iteration-count", "1",
				"--output-path", ".png", "--threads", "5", "--seed", "1", "--affine-params", "1,1,1,1,1,1", "--functions", "swirl:0.5,horseshoe:0.1",
				"--config", s.configFile.Name(), "--gamma", "2", "-s", "100",
			},
			expectedArgs: &domain.Args{
				Size:            domain.Size{Height: 1, Width: 1},
				IterationCount:  1,
				OutputPath:      ".png",
				Threads:         5,
				Seed:            1,
				Functions:       []domain.Function{{Name: domain.Swirl, Weight: 0.5}, {Name: domain.Horseshoe, Weight: 0.1}},
				AffineParams:    []domain.AffineParam{{A: 1, B: 1, C: 1, D: 1, E: 1, F: 1}},
				GammaCorrection: true,
				Gamma:           2,
				SymmetryLevel:   100,
			},
		},
		{
			name: "only config",
			args: []string{"fractal-flame", "--config", s.configFile.Name()},
			expectedArgs: &domain.Args{
				Size:            domain.Size{Height: 999, Width: 999},
				IterationCount:  999,
				OutputPath:      "cfg.png",
				Threads:         9,
				Seed:            9.99,
				Functions:       []domain.Function{{Name: domain.Swirl, Weight: 0.5}, {Name: domain.Horseshoe, Weight: 0.2}},
				AffineParams:    []domain.AffineParam{{A: 2, B: 2, C: 2, D: 2, E: 2, F: 2}},
				GammaCorrection: true,
				Gamma:           2,
				SymmetryLevel:   5,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.service.EXPECT().RenderFlame(tt.expectedArgs).Return(nil)
			err := s.app.Run(context.Background(), tt.args)
			s.Require().NoError(err)
		})
	}
}

func (s *appSuite) TestServiceReturnErr() {
	s.Run("service return err", func() {
		s.service.EXPECT().RenderFlame(&domain.Args{
			Size:            domain.Size{Height: 1080, Width: 1920},
			IterationCount:  2500,
			OutputPath:      "result.png",
			Threads:         1,
			Seed:            5.1234,
			Functions:       []domain.Function{{Name: domain.Swirl, Weight: 1.0}},
			AffineParams:    []domain.AffineParam{{A: 0.9, B: 0.7, C: 0, D: -0.15, E: -1.1, F: 0}},
			GammaCorrection: false,
			Gamma:           2.2,
			SymmetryLevel:   1,
		}).Return(fmt.Errorf("some error"))
		err := s.app.Run(context.Background(), []string{"fractal-flame"})
		s.Require().Error(err)
	})
}

func (s *appSuite) TestParseArgsErrors() {
	tests := []struct {
		name        string
		args        []string
		expectedErr string
	}{
		{
			name:        "wrong config file",
			args:        []string{"fractal-flame", "--config", s.wrongFile.Name()},
			expectedErr: "cannot parse config file",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.app.Run(context.Background(), tt.args)
			s.Require().ErrorContains(err, tt.expectedErr)
		})
	}
}

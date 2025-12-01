package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type configSuite struct {
	suite.Suite
	app             *App
	wrongConfigFile *os.File
	logBuffer       io.Writer
}

func TestRunConfigSuite(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (s *configSuite) SetupTest() {
	s.logBuffer = &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(s.logBuffer, &slog.HandlerOptions{Level: slog.LevelError}))
	s.app = NewApp(nil, logger)

	tempDir := s.T().TempDir()
	var err error

	s.wrongConfigFile, err = os.CreateTemp(tempDir, "wrong-*.json")
	s.Require().NoError(err)
}

func (s *configSuite) SetConfigContent(cfgData domain.Args) {
	jsonBytes, err := json.Marshal(cfgData)
	s.Require().NoError(err)
	_, err = s.wrongConfigFile.Write(jsonBytes)
	s.Require().NoError(err, err)
}

func (s *configSuite) TestParseWrongDimension() {
	cfg := domain.Args{
		Size: domain.Size{Height: -1, Width: 999},
	}
	s.SetConfigContent(cfg)

	s.Run("wrong dimension", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongItCount() {
	cfg := domain.Args{
		IterationCount: -20,
	}
	s.SetConfigContent(cfg)

	s.Run("wrong iteration count", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongOutputPath() {
	cfg := domain.Args{
		OutputPath: "123.txt",
	}
	s.SetConfigContent(cfg)

	s.Run("wrong output path", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongAmountOfThreads() {
	cfg := domain.Args{
		Threads: -20,
	}
	s.SetConfigContent(cfg)

	s.Run("wrong amount of threads", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongFunctionName() {
	cfg := domain.Args{
		Functions: []domain.Function{{Name: "wrong name", Weight: 0.5}},
	}
	s.SetConfigContent(cfg)

	s.Run("wrong function name", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongfunctionWeight() {
	cfg := domain.Args{
		Functions: []domain.Function{{Name: domain.Swirl, Weight: -20}},
	}
	s.SetConfigContent(cfg)

	s.Run("wrong function weight", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseGammaEqZero() {
	cfg := domain.Args{
		Gamma: 0,
	}
	s.SetConfigContent(cfg)

	s.Run("gamma value equal zero", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().NoError(err)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestAffines() {
	cfg := domain.Args{
		AffineParams: []domain.AffineParam{{A: 0, B: 0, C: 0, D: 0, E: 0, F: 0}, {A: 1, B: 1, C: 1, D: 1, E: 1, F: 1}},
	}
	s.SetConfigContent(cfg)

	args := &domain.Args{}

	s.Run("must skip param with all zeros", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, args)
		s.Require().NoError(err)
		s.Require().Equal([]domain.AffineParam{{A: 1, B: 1, C: 1, D: 1, E: 1, F: 1}}, args.AffineParams)
	})

	_ = s.wrongConfigFile.Close()
}

func (s *configSuite) TestParseWrongSymmetryLevel() {
	cfg := domain.Args{
		SymmetryLevel: -20,
	}
	s.SetConfigContent(cfg)

	s.Run("wrong amount of symmetry level", func() {
		err := s.app.readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errWrongSymmetryLevel)
	})

	_ = s.wrongConfigFile.Close()
}

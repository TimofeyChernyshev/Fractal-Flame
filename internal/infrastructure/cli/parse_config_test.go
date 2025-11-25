package cli

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v3"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
)

type configSuite struct {
	suite.Suite
	wrongConfigFile *os.File
}

func TestRunConfigSuite(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (s *configSuite) SetupTest() {
	tempDir := s.T().TempDir()
	var err error

	s.wrongConfigFile, err = os.CreateTemp(tempDir, "wrong-*.json")
	s.Require().NoError(err)
}

func (s *configSuite) SetConfigContent(cfgData domain.Args) {
	jsonBytes, err := json.Marshal(cfgData)
	s.Require().NoError(err)
	_, err = s.wrongConfigFile.Write(jsonBytes)
	s.Require().NoError(err)
}

func (s *configSuite) TestParseWrongDimension() {
	cfg := domain.Args{
		Size: domain.Size{Height: -1, Width: 999},
	}
	s.SetConfigContent(cfg)

	s.Run("wrong dimension", func() {
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
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
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
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
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
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
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
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
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
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
		err := readConfig(s.wrongConfigFile.Name(), &cli.Command{}, &domain.Args{})
		s.Require().ErrorIs(err, errConfigArgs)
	})

	_ = s.wrongConfigFile.Close()
}

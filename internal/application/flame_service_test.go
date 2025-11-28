package application

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"math"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

type serviceSuite struct {
	suite.Suite
	saver     *MockSaver
	renderer  *MockRenderer
	chooser   *MockChooser
	randomGen *MockRandomGenerator
	random    *random.MockRandom
	ctrl      *gomock.Controller

	service   *FlameService
	logBuffer io.Writer
	args      *domain.Args
	image     *domain.FractalImage
}

func TestRunAppSuite(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}

func (s *serviceSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.saver = NewMockSaver(s.ctrl)
	s.renderer = NewMockRenderer(s.ctrl)
	s.chooser = NewMockChooser(s.ctrl)
	s.randomGen = NewMockRandomGenerator(s.ctrl)
	s.random = random.NewMockRandom(s.ctrl)

	s.logBuffer = &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(s.logBuffer, &slog.HandlerOptions{Level: slog.LevelError}))

	s.service = NewFlameService(s.saver, s.chooser, s.randomGen, logger)

	s.args = &domain.Args{
		Size:           domain.Size{Height: 999, Width: 999},
		IterationCount: 999,
		OutputPath:     "cfg.png",
		Threads:        1,
		Seed:           9.99,
		Functions: []domain.Function{
			{Name: domain.Swirl, Weight: 0.5},
			{Name: domain.Horseshoe, Weight: 0.2},
		},
		AffineParams: domain.AffineParam{
			A: 2, B: 2, C: 2, D: 2, E: 2, F: 2,
		},
	}

	s.image = domain.NewFractalImage(999, 999)
}

func (s *serviceSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *serviceSuite) TestParseArgs() {
	seed := int64(math.Float64bits(s.args.Seed))

	s.Run("No errors", func() {
		s.randomGen.EXPECT().New(seed).Return(s.random)
		s.chooser.EXPECT().Choose(s.args.Threads, s.random).Return(s.renderer)
		s.renderer.EXPECT().Render(s.args).Return(s.image)
		s.saver.EXPECT().Save(s.image, "cfg.png").Return(nil)

		err := s.service.RenderFlame(s.args)
		s.Require().NoError(err)
	})
}

func (s *serviceSuite) TestSaverReturnErr() {
	seed := int64(math.Float64bits(s.args.Seed))

	s.Run("saver return error", func() {
		s.randomGen.EXPECT().New(seed).Return(s.random)
		s.chooser.EXPECT().Choose(s.args.Threads, s.random).Return(s.renderer)
		s.renderer.EXPECT().Render(s.args).Return(s.image)
		s.saver.EXPECT().Save(s.image, "cfg.png").Return(fmt.Errorf("some error"))

		err := s.service.RenderFlame(s.args)
		s.Require().Error(err)
	})
}

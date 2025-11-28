package renderers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw4-fractal-flame/pkg/random"
)

func TestChooserChoose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRnd := random.NewMockRandom(ctrl)

	chooser := NewChooser()

	t.Run("single thread returns SingleThreadRenderer", func(t *testing.T) {
		r := chooser.Choose(1, mockRnd)
		if _, ok := r.(*SingleThreadRenderer); !ok {
			t.Fatalf("expected SingleThreadRenderer, got %T", r)
		}
	})

	t.Run("multi thread returns MultiThreadRenderer", func(t *testing.T) {
		r := chooser.Choose(4, mockRnd)
		if _, ok := r.(*MultiThreadRenderer); !ok {
			t.Fatalf("expected MultiThreadRenderer, got %T", r)
		}
	})
}

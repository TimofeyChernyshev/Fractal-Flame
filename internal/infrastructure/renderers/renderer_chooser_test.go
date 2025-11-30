package renderers

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestChooserChoose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRndGen := NewMockRandomGenerator(ctrl)

	chooser := NewChooser(mockRndGen)

	t.Run("single thread returns SingleThreadRenderer", func(t *testing.T) {
		r := chooser.Choose(1)
		if _, ok := r.(*SingleThreadRenderer); !ok {
			t.Fatalf("expected SingleThreadRenderer, got %T", r)
		}
	})

	t.Run("multi thread returns MultiThreadRenderer", func(t *testing.T) {
		r := chooser.Choose(4)
		if _, ok := r.(*MultiThreadRenderer); !ok {
			t.Fatalf("expected MultiThreadRenderer, got %T", r)
		}
	})
}

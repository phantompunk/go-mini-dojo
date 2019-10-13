package shape

import (
	"testing"
)

func TestVolume(t *testing.T) {
	checkVolume := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Volume()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("find volume of a cube", func(t *testing.T) {
		cube := Cube{4.0}
		checkVolume(t, cube, 64.0)
	})

	t.Run("find volume of rectangular box", func(t *testing.T) {
		box := RectangularBox{
			Length: 4.0,
			Width:  3.0,
			Height: 5.0,
		}
		checkVolume(t, box, 60.0)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("find area of a cube", func(t *testing.T) {
		cube := Cube{6.0}
		checkArea(t, cube, 216.0)
	})

	t.Run("find area of a rectangular box", func(t *testing.T) {
		box := RectangularBox{
			Length: 4.0,
			Width:  3.0,
			Height: 5.0,
		}
		checkArea(t, box, 94.0)
	})
}

func TestShapeVolume(t *testing.T) {
	tests := []struct {
		shape Shape
		want  float64
	}{
		{Cube{4.0}, 64.0},
		{RectangularBox{4.0, 3.0, 5.0}, 60.0},
	}
	for _, tt := range tests {
		s := tt.shape
		if got := s.Volume(); got != tt.want {
			t.Errorf("Shape.Volume() = %v, want %v", got, tt.want)
		}
	}
}

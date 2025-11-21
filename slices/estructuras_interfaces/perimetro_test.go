package estructura

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimetro(rect)

	want := 40.0

	if got != want {
		t.Errorf("te dio %.2f y te debio dar %.2f", got, want)
	}

}
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 36.0},
		{circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %#v want %#v", got, tt.want)
		}
	}

}

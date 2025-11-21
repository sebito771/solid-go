package iteracion

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")

	expected := "aaaaa"

	if repeated != expected {

		t.Errorf("repeated '%q' pero tenemos '%q' ", expected, repeated)

	}

}

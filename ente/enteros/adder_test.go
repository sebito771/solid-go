package intenger

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) { //RECUERDA QUE DESPUES DEL TEST LA PALABRA DEBE EMPEZAR EN MAYUSCULAS
	sum := Add(2, 2)

	expected := 4

	if sum != expected {
		t.Errorf("expected '%d'but got '%d'", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output : 6
}

package main

import (
	"bytes"
	"testing"
)

func TestCuenta(t *testing.T) {
	buffer := &bytes.Buffer{}

	Cuenta(buffer)

	got := buffer.String()
	want := `3
	2
	1
	go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

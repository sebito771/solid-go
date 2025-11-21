package main

import (
	"bytes"
	"testing"
)

func TestSaludo(t *testing.T) {
	buffer := bytes.Buffer{}
	Saludo(&buffer, "chris")

	got := buffer.String()
	want := "Hello, chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

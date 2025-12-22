package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Saludo(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)

}

func SaludoHandler(w http.ResponseWriter, r *http.Request) {
	Saludo(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(SaludoHandler)))
}
// me voy a comprometer a aprender este leenguaje de programacion

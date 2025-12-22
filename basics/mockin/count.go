package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	iniciaC   = 3
	FinalWord = "go!"
)

func Cuenta(out io.Writer) {
	for i := iniciaC; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Minute)
	}
	fmt.Fprint(out, FinalWord)

}

func main() {
	Cuenta(os.Stdout)
}

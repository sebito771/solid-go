package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64
	var operacion string
	fmt.Print("dame un numero: ")
	fmt.Scanln(&numero1)
	fmt.Print("dame otro numero: ")
	fmt.Scanln(&numero2)
	fmt.Print("hacer suma(+),resta(-),multiplicacion(*),division(/): ")
	fmt.Scanln(&operacion)

	switch operacion {
	case "+":
		suma := numero1 + numero2
		fmt.Println("suma", suma)

	case "-":
		resta := numero1 - numero2
		fmt.Println("resta", resta)

	case "*":
		multiplicacion := numero1 * numero2
		fmt.Println("multiplicacion", multiplicacion)

	case "/":
		if numero2 != 0 {
			division := numero1 / numero2
			fmt.Println("división:", division)
		} else {
			fmt.Println("No se puede dividir entre 0")
		}
	}
}

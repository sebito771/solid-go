package main

import "fmt"

func calcular(n1 float64, n2 float64, op string) float64 {
	var resultado float64
	switch op {
	case "+":
		resultado = n1 + n2
	
	case "-":
		resultado= n1-n2
	
	case "*":
		resultado= n1*n2
	
	case "/":
		resultado= n1/n2
	
	default:
		fmt.Println("hola guapo, pasaste mal los datoss")
	}
	return resultado

}


func main(){
		var numero1 float64
	var numero2 float64
	var operacion string
	fmt.Print("dame un numero: ")
	fmt.Scanln(&numero1)
	fmt.Print("dame otro numero: ")
	fmt.Scanln(&numero2)
	fmt.Print("hacer suma(+),resta(-),multiplicacion(*),division(/): ")
	fmt.Scanln(&operacion)


	resultado:= calcular(numero1,numero2,operacion)
	fmt.Print("resultado: ",resultado)
}
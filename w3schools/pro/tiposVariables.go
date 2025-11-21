package main

import (
	"fmt"	
)

/*Declara un entero y un decimal.
Suma ambos (haz la conversión que se necesite) y muestra el resultado.
Convierte el decimal a entero y muestra el resultado truncado.
*/


func main(){
	var entero int = 2
	var decimal float64= 3.0
	sum := float64(entero)+ decimal
fmt.Println("hola guapo")
fmt.Printf("estos son los valores sumados %v , %.2f, y este es el resultado %.2f",entero,decimal,sum)
}


package main

import "fmt"


func calificar_num(n1 float64) string{
	var resultado string
	if n1>=1 {
		resultado= "positivo"	
	}
	if n1==0{
      resultado= "cero"
	}
	if n1<0{
		resultado= "negativo"
	}
	return	resultado
}

func main (){
	fmt.Println("digame un numero")
	var numero float64
	fmt.Scanln(&numero)
	resultado:= calificar_num(numero)
	fmt.Println("el numero que diste es ",resultado)
}


package main

import (
	"fmt"
	
)

func calificarArray(numeros []float64) (cero, positivos, negativos int) {

     for _,num:= range numeros{
		if num==0{
			cero++
		}else if num<0{
			negativos++
		}else{
			positivos++
		}
	
 }
 return cero,positivos,negativos
}

func main() {
	fmt.Println("ingrese una lista de numeros")
	var numero float64
    lista:= make([]float64,0)
     terminar:= ""
    for terminar!="x" {
         fmt.Println("digame uun numero")
		 fmt.Scanln(&numero)
		 lista = append(lista, numero)
		 fmt.Println("lista=",lista)
		 fmt.Println("desea continuar? x para salir")
		 fmt.Scanln(&terminar)
	}
 cero,positivos,negativos:= calificarArray(lista)
fmt.Println("cero",cero)
fmt.Println("negativos",negativos)
fmt.Println("positivos",positivos)
}
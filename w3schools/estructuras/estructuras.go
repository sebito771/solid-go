package main

import "fmt"

type Persona struct {
	nombre string
	edad   int

}

func calificarEdad(list []Persona)(mayor,menor []string ){
	var respuesta []string
	var respuestaMenor []string
 for _,ed:= range list{
	if ed.edad>=18{
		respuesta = append(respuesta, ed.nombre)
	}else{
		respuestaMenor= append(respuestaMenor, ed.nombre)
	}
 }
 return respuesta,respuestaMenor
}


func main() {
	var num_personas int
    fmt.Println("cuantas personnas desea añadir")
	fmt.Scanln(&num_personas)
	var person Persona
    var list []Persona

	for i:=0;i<num_personas;i++{
      fmt.Println("dime el nombre de la persona")
	  fmt.Scanln(&person.nombre)
	  fmt.Println("dime la edad de la persona")
	  fmt.Scanln(&person.edad)
	  list = append(list, person)
	 // fmt.Println(list)
	}
	Mayor,menor:= calificarEdad(list)
	fmt.Println("estos son los mayores de edad",Mayor)
	fmt.Println("esta es la cantidad de mayores de edad",len(Mayor))
	fmt.Println("estos son los menores de edad",menor)
	fmt.Println("esta es la cantidad de menores de edad",len((menor)))
}
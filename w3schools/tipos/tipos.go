//Declara cuatro variables: un nombre (string), una edad (int), un promedio (float64), y un estado (bool).
//Imprime cada una y muestra su tipo de dato.

package main
import "fmt"
import "reflect"

func main(){
    var nombre string
	var edad int
	var promedio float64
	var estado bool
	fmt.Println(reflect.TypeOf(nombre),nombre)
	fmt.Println(reflect.TypeOf(edad),edad)
    fmt.Println(reflect.TypeOf(promedio,),promedio)
	fmt.Println(reflect.TypeOf(estado),estado)
}


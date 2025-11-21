package main

import (
	"log"
	"net/http"
)

//tipos y demas
type InMemoryPlaystore struct{
	score map[string]int
}
func NewInMemoryPlaystore () *InMemoryPlaystore{
		return &InMemoryPlaystore{
		score: make(map[string]int),
	}
}

//funciones
func(i *InMemoryPlaystore)Puntaje(name string)int{
	return i.score[name]
}
func(i *InMemoryPlaystore)RegistrarVictoria(name string){
    i.score[name]++
}

func main() {
	store:= NewInMemoryPlaystore()
	server := &playerServer{store}
	log.Fatal(http.ListenAndServe(":5000", server))
}

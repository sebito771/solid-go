package main

import (
	"fmt"
	"net/http"
	"strings"

	
)

//tipos


type Playerstore interface {
	Puntaje(name string) int
	RegistrarVictoria(name string)
	
}
type playerServer struct {
	store Playerstore
}
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}


// funciones


func (p *playerServer) showScore(w http.ResponseWriter, r *http.Request) {
	
	
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score:= p.store.Puntaje(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,score)
}
func (p *playerServer)ServeHTTP(w http.ResponseWriter,r *http.Request){
	player:= strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
		
	case http.MethodPost:
		p.processWin(w,player)
	case http.MethodGet:
		p.showScore(w, r)
	}

}
func(p *playerServer)processWin(w http.ResponseWriter, player string){
	p.store.RegistrarVictoria(player)
	w.WriteHeader(http.StatusAccepted)
}




package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//interfaces,estructuras, tipos en general

type StubPlayerStore struct {
	scores map[string]int

}
 // funciones
 func (s *StubPlayerStore) Puntaje(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RegistrarVictoria(name string) {
    s.scores[name]++
}

func TestGetPlayer(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"sebas": 20,
			"chori": 100,
		},

	    
	}
    server := playerServer{&store}
	t.Run("puntaje sebas", func(t *testing.T) {
		request := (NewScoreRequest("sebas"))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("puntaje chori", func(t *testing.T) {
		request := NewScoreRequest("chori")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "100")
	})

	t.Run("el jugador no es encontrado", func(t *testing.T) {
		request := NewScoreRequest("apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)

	})

}
func TestStoreWins(t *testing.T){
	store:= StubPlayerStore{
		map[string]int{},

	}
	server:= playerServer{&store}

t.Run("returns el post",func(t *testing.T) {
	request,_:= http.NewRequest(http.MethodPost,"/players/sebas",nil)
	response:= httptest.NewRecorder()
	server.ServeHTTP(response,request)
	assertStatus(t,response.Code,http.StatusAccepted)
})
}

// funnciones auxiliares
func NewScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func AssertResponseBody(t testing.TB, got, want string) error {
	t.Helper()
	if got != want {
		t.Errorf("hay algo mal con el Body de Response,got %q want %q", got, want)
	}
	return nil
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("estatus incorrecto %d want %d", got, want)
	}
}


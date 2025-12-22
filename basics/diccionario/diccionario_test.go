package main

import (
	"testing"
)

// PRUEBAS
func TestDicc(t *testing.T) {
	diccionari := diccionario{"test": "esto es una prueba"}
	t.Run("palabras conocidas", func(t *testing.T) {
		got, _ := diccionari.Search("encontrada")
		want := " no se ha encontrado la palabra con exito"

		acierto(t, got, want)

	})
	t.Run("palabras desconocidas", func(t *testing.T) {
		_, got := diccionari.Search("desconocido")
		if got == nil {
			t.Fatal("se espera tener un error")
		}
		assertError(t, got, ErrorNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("palabra nueva", func(t *testing.T) {
		dicc := diccionario{}
		word := "test"
		definition := "esto solo es una prueba"
		err := dicc.Add(word, definition)
		assertError(t, err, nil)
		asssertDefinition(t, dicc, word, definition)
	})

	t.Run("palabras existentes", func(t *testing.T) {
		word := "prueba"
		definition := "esto es una prueba"
		dicci := diccionario{word: definition}
		err := dicci.Add(word, "new test")
		assertError(t, err, ErrorWordExist)
		asssertDefinition(t, dicci, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("palabras existentes", func(t *testing.T) {
		word := "test"
		definition := "esto es un test"
		dicc := diccionario{word: definition}
		newDefinition := "nueva definicion"
		err := dicc.Update(word, newDefinition)
		assertError(t, err, nil)
		asssertDefinition(t, dicc, word, newDefinition)
	})

	t.Run("nueva palabra", func(t *testing.T) {
		word := "test"
		definition := "esto es una prueba"
		dicc := diccionario{}
		err := dicc.Update(word, definition)
		assertError(t, err, ErrorWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("palabra existentes", func(t *testing.T) {
		word := "test"
		dicc := diccionario{word: "test definition"}
		err := dicc.Delete(word)
		assertError(t, err, nil)

		_, err = dicc.Search(word)

		assertError(t, err, ErrorNotFound)

	})

	t.Run("palabra no existente", func(t *testing.T) {
		word := "test"
		dic := diccionario{}
		err := dic.Delete(word)
		assertError(t, err, ErrorWordDoesNotExist)
	})
}

// AYUDANTES
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("tienes un error %q cuando esperabas %q", got, want)
	}
}

func acierto(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func asssertDefinition(t testing.TB, dicc diccionario, word, definition string) {
	t.Helper()
	got, err := dicc.Search(word)
	if err != nil {
		t.Fatal("se deberia encontrar la palabra agregada", err)
	}
	acierto(t, got, definition)
}

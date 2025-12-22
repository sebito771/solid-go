package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sumar slice dinámico", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
   
		if got != want {
			t.Errorf("te dio %d, te debio dar %d, y tus datos son %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := Sumall([]int{1, 2}, []int{0, 9})
	resultado := []int{3, 9}

	if !reflect.DeepEqual(got, resultado) {
		t.Errorf("got %v want %v", got, resultado)
	}
}

func TestSumalltails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := Sumalltails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := Sumalltails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

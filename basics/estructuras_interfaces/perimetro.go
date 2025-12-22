package estructura

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   float64
	altura float64
}

type Rectangle struct {
	altura float64
	ancho  float64
}

type circle struct {
	radio float64
}

func Perimetro(rect Rectangle) float64 {
	return 2 * (rect.altura + rect.ancho)

}

func Area(rect Rectangle) float64 {

	return (rect.altura * rect.ancho)
}

func (rect Rectangle) Area() float64 {
	return 2 * (rect.altura + rect.ancho)

}

func (c circle) Area() float64 {
	return math.Pi * c.radio * c.radio
}

func (t Triangle) Area() float64 {
	return (t.base * t.altura) * 0.5
}

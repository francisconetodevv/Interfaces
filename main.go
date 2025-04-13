package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64 /* Método chamado área */
}

type retangulo struct {
	largura, altura float64
}

/* O retangulo está implementado a interface chamada geometria */
func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	radius float64
}

/* O circulo está implementado a interface chamada geometria */
func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())
	/* Através essa interface, precisa implementar esse método chamado área*/
}

func main() {
	fmt.Println("Inicializando projeto...")

	retangulo := retangulo{
		largura: 1,
		altura:  2,
	}

	circulo := circulo{
		radius: 4,
	}

	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)
}

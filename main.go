package main

import (
	"errors"
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

	p := ProblemaNetwork{
		rede:     true,
		hardware: false,
	}

	ExibirError(errors.New("A error"))
	ExibirError(p)
}

type ProblemaNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaNetwork) Error() string {
	if p.rede {
		return "Problema de Rede"
	} else if p.hardware {
		return "Problema de Hardware"
	} else {
		return "Outro problema"
	}
}

func ExibirError(err error) {
	fmt.Println(err.Error())
}

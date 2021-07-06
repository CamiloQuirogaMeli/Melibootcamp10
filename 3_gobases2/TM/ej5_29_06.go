package main

import (
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(a string) func(float64) float64 {
	switch a {
	case perro:
		return perroFunc
	case gato:
		return gatoFunc
	case tarantula:
		return tarantulaFunc
	case hamster:
		return hamsterFunc
	}
	return nil
}

func perroFunc(valor float64) float64 {
	var cantidad float64 = valor * 10
	return cantidad
}

func gatoFunc(valor float64) float64 {
	var cantidad float64 = valor * 5
	return cantidad
}

func hamsterFunc(valor float64) float64 {
	var cantidad float64 = valor * 0.25
	return cantidad
}

func tarantulaFunc(valor float64) float64 {
	var cantidad float64 = valor * 0.15
	return cantidad
}

func main() {

	animalPerro := Animal(perro)
	animalGato := Animal(gato)
	animalTarantula := Animal(tarantula)
	animalHamster := Animal(hamster)

	var cantidad float64

	cantidad += animalPerro(5)
	cantidad += animalGato(5)
	cantidad += animalTarantula(5)
	cantidad += animalHamster(5)

	fmt.Println("cantidad de alimento: ", cantidad)

}

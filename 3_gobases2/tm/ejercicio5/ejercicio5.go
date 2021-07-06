package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {

	animalPerro, err := animal(perro)
	animalGato, err := animal(gato)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	if err != nil {
		fmt.Println(err)
	} else {
		cantidad := 0.0
		cantidad += animalPerro(2)
		cantidad += animalGato(2)
		cantidad += animalTarantula(2)
		cantidad += animalHamster(2)

		fmt.Println(cantidad)
	}
}

func animal(op string) (func(value int) float64, error) {
	switch op {
	case perro:
		return perroFunc, nil
	case gato:
		return gatoFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	}
	return nil, errors.New("Error")
}

func perroFunc(value int) float64 {
	return float64(value * 10)
}

func gatoFunc(value int) float64 {
	return float64(value * 5)
}

func hamsterFunc(value int) float64 {
	return float64(value) * float64(25) / float64(100)
}

func tarantulaFunc(value int) float64 {
	return float64(value) * float64(15) / float64(100)
}

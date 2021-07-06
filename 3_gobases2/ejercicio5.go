package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func ejercicio5() {
	animalPerro, err := Animal(PERRO)
	if err != nil {
		fmt.Println(err)
	}
	animalGato, err := Animal(GATO)
	if err != nil {
		fmt.Println(err)
	}
	animalHamster, err := Animal(HAMSTER)
	if err != nil {
		fmt.Println(err)
	}
	animalTarantula, err := Animal(TARANTULA)
	if err != nil {
		fmt.Println(err)
	}

	var cantidad float64
	cantidad += animalPerro(5)
	fmt.Println(cantidad)
	cantidad += animalGato(8)
	fmt.Println(cantidad)
	cantidad += animalHamster(3)
	fmt.Println(cantidad)
	cantidad += animalTarantula(9)
	fmt.Println(cantidad)
}

func Animal(tipoAnimal string) (func(cantidad float64) float64, error) {

	switch tipoAnimal {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTER:
		return animalHamster, nil
	case TARANTULA:
		return animalTarantula, nil
	default:
		return nil, errors.New("no existe definicion para el animal")
	}

}

func animalPerro(cantidad float64) float64 {
	return cantidad * 10
}

func animalGato(cantidad float64) float64 {
	return cantidad * 5
}

func animalHamster(cantidad float64) float64 {
	return cantidad * 0.25
}

func animalTarantula(cantidad float64) float64 {
	return cantidad * 0.15
}

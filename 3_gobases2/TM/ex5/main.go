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

func Animal(name string) (func(cant int) float64, error) {
	switch name {
	case PERRO:
		return perroComida, nil
	case GATO:
		return gatoComida, nil
	case HAMSTER:
		return hamsterComida, nil
	case TARANTULA:
		return tarantulaComida, nil
	}
	return nil, errors.New("inst invalida")
}

func perroComida(cant int) float64 {
	return float64(cant) * float64(10000)
}

func gatoComida(cant int) float64 {
	return float64(cant) * float64(5000)
}

func hamsterComida(cant int) float64 {
	return float64(cant) * float64(250)
}

func tarantulaComida(cant int) float64 {
	return float64(cant) * float64(150)
}

func main() {
	animalPerro, err := Animal(PERRO)
	animalGato, err := Animal(GATO)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)

	var cantidad float64
	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(20)
	cantidad += animalTarantula(25)

	if err == nil {
		fmt.Println("Cantidad: ", cantidad)
	}
}

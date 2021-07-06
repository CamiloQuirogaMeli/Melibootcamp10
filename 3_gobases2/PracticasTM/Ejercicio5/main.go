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

func Animal(animal string) (func(cantidad int) int, error) {
	switch animal {
	case PERRO:
		return perro, nil
	case GATO:
		return gato, nil
	case HAMSTER:
		return hamster, nil
	case TARANTULA:
		return tarantula, nil
	default:
		return nil, errors.New("El animal seleccionado no existe")
	}
}

func perro(cantidad int) int {
	cantidad *= 10000
	return cantidad
}

func gato(cantidad int) int {
	cantidad *= 5000
	return cantidad
}

func hamster(cantidad int) int {
	cantidad *= 250
	return cantidad
}

func tarantula(cantidad int) int {
	cantidad *= 150
	return cantidad
}

func main() {
	animalPerro, err := Animal(PERRO)
	animalGato, err := Animal(GATO)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)

	var cantidad int
	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(3)
	cantidad += animalTarantula(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Necesitas comprar ", cantidad, "gramos de comida en total")
	}
}

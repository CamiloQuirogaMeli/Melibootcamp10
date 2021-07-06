package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTERS  = "hamster"
	TARANTULA = "tarantula"
)

func animalPerro(cant float64) float64 {
	return cant * 10
}

func animalGato(cant float64) float64 {
	return cant * 5
}

func animalHamster(cant float64) float64 {
	return cant * 0.25
}

func animalTarantula(cant float64) float64 {
	return cant * 0.15
}

func Animal(animal string) (func(cant float64) float64, error) {
	switch animal {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTERS:
		return animalHamster, nil
	case TARANTULA:
		return animalTarantula, nil

	}
	return nil, errors.New("El animal no existe.")
}

func main() {
	var cantidad float64
	animalPerro, err := Animal(PERRO)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalPerro(5)
		fmt.Print("Alimento Perro: ", cantidad, "\n")
	}

	animalGato, err := Animal(GATO)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalGato(8)
		fmt.Print("Alimento Gato: ", cantidad, "\n")
	}

	animalHamster, err := Animal(HAMSTERS)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalHamster(2)
		fmt.Print("Alimento Hamster: ", cantidad, "\n")
	}

	animalTarantula, err := Animal(TARANTULA)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalTarantula(3)
		fmt.Print("Alimento Tarantula: ", cantidad, "\n")
	}
}

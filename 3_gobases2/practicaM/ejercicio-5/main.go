package main

import (
	"errors"
	"fmt"
)

func main() {
	funcAnimalPerro, err := Animal(perro)
	funcAnimalGato, err := Animal(gato)
	funcAnimalHamster, err := Animal(hamster)
	funcAnimalTarantula, err := Animal(tarantula)

	if err == nil {
		fmt.Println(funcAnimalPerro(50), " kg")
		fmt.Println(funcAnimalGato(100), " kg")
		fmt.Println(funcAnimalHamster(500), " grms")
		fmt.Println(funcAnimalTarantula(1000), "grms")
	} else {
		fmt.Println(err)
	}
}

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(tipoAnimal string) (func(cantAnimal int) float64, error) {
	switch tipoAnimal {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	}
	return nil, errors.New("ERROR! El animal no existe")
}

func animalPerro(cantAnimal int) float64 {
	const kg float64 = 10
	return float64(cantAnimal) * kg
}

func animalGato(cantAnimal int) float64 {
	const kg float64 = 5
	return float64(cantAnimal) * kg
}

func animalHamster(cantAnimal int) float64 {
	const grms float64 = 250
	return float64(cantAnimal) * grms
}

func animalTarantula(cantAnimal int) float64 {
	const grms float64 = 150
	return float64(cantAnimal) * grms
}

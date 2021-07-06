package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case "perro":
		return alimentoPerro, nil
	case "gato":
		return alimentoGato, nil
	case "hamster":
		return alimentoHamster, nil
	case "tarantula":
		return alimentoTarantula, nil
	default:
		return nil, errors.New("ingrese un animal valido")
	}
}

func alimentoPerro(perros int) float64 {
	return 10000 * float64(perros)
}

func alimentoGato(gatos int) float64 {
	return 5000 * float64(gatos)
}

func alimentoTarantula(tarantulas int) float64 {
	return 150 * float64(tarantulas)
}

func alimentoHamster(hamsters int) float64 {
	return 250 * float64(hamsters)
}

func main() {
	var cantidad float64
	animalPerro, err := Animal(perro)

	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalPerro(5)
	}

	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalGato(3)
	}

	animalHamster, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalHamster(2)
	}

	animalTarantula, err := Animal(tarantula)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalTarantula(3)
	}

	fmt.Println("Cantidad total de alimento:", cantidad/1000, "kg")
}

package main

import (
	"errors"
	"fmt"
)

const (
	DOG            = "dog"
	CAT            = "cat"
	HAMSTER        = "hamster"
	TARANTULA      = "tarantula"
	DOG_FOOD       = 10
	CAT_FOOD       = 5
	HAMSTER_FOOD   = 250
	TARANTULA_FOOD = 150
)

func dogFunc(animals int) float64 {
	return float64(animals * DOG_FOOD)
}
func catFunc(animals int) float64 {
	return float64(animals * CAT_FOOD)
}
func hamsterFunc(animals int) float64 {
	return float64(animals * HAMSTER_FOOD)
}
func tarantulaFunc(animals int) float64 {
	return float64(animals * TARANTULA_FOOD)
}

func Animal(animalType string) (func(int) float64, error) {
	switch animalType {
	case DOG:
		return dogFunc, nil
	case CAT:
		return catFunc, nil
	case HAMSTER:
		return hamsterFunc, nil
	case TARANTULA:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("No se pudo calcular ninguna quantity")
	}
}

func main() {

	animalDog, err := Animal(DOG)
	animalCat, err := Animal(CAT)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)
	var quantity float64

	if err != nil {
		fmt.Println(err)
	} else {
		quantity += animalDog(5)
		quantity += animalCat(8)
		quantity += animalHamster(12)
		quantity += animalTarantula(7)
		fmt.Printf("La quantity en kg a comprar de comida es %.2f\n", quantity)
	}
}

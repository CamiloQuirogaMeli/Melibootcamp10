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

func calculateDogFood(totalAnimals int) float64 {
	return float64(totalAnimals * 10)
}

func calculateCatFood(totalAnimals int) float64 {
	return float64(totalAnimals * 5)
}

func calculateHamsterFood(totalAnimals int) float64 {
	totalFood := float64(totalAnimals*250) / 1000
	return totalFood
}

func calculateScorpionFood(totalAnimals int) float64 {
	totalFood := float64(totalAnimals*150) / 1000
	return totalFood
}

func Animal(animal string) (func(totalAnimals int) float64, error) {

	switch animal {
	case PERRO:
		return calculateDogFood, nil
	case GATO:
		return calculateCatFood, nil
	case HAMSTER:
		return calculateHamsterFood, nil
	case TARANTULA:
		return calculateScorpionFood, nil
	}

	return nil, errors.New("The animal received is not a valid animal to calculate food")
}

func main() {
	var cantidad float64

	animalPerro, err := Animal(PERRO)
	animalGato, err := Animal(GATO)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)

	if err != nil {
		fmt.Println(err)
	}

	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(2)
	cantidad += animalTarantula(3)

	fmt.Printf("The total food to buy is %f", cantidad)
}

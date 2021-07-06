package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     string = "perro"
	GATO      string = "gato"
	HAMSTER   string = "hamster"
	TARANTULA string = "tarantula"
)

func main() {
	function, err := Animal(TARANTULA)
	if err != nil {
		fmt.Println(err)
	} else {
		result := function(5)
		fmt.Println(result)
	}
}

func Animal(animal string) (func(val int) float64, error) {
	switch animal {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTER:
		return animalHamster, nil
	case TARANTULA:
		return animalTarantula, nil
	default:
		return nil, errors.New("El animal aun no esta definido")
	}
}

func animalPerro(amount int) float64 {
	return float64(amount) * 10
}
func animalGato(amount int) float64 {
	return float64(amount) * 5
}
func animalHamster(amount int) float64 {
	return float64(amount) * 0.25
}
func animalTarantula(amount int) float64 {
	return float64(amount) * 0.15
}

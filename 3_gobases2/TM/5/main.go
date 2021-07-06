package main

import (
	"errors"
	"fmt"
)

func main() {

	const (
		PERRO     = "perro"
		GATO      = "gato"
		HAMSTER   = "hamster"
		TARANTULA = "tarantula"
	)

	animalPerro, err := animal(PERRO)
	animalGato, err2 := animal(GATO)
	animalHamster, err3 := animal(HAMSTER)
	animalTarantula, err4 := animal(TARANTULA)

	var cantidad float64

	if err == nil {
		cantidad += animalPerro(5)
	} else {
		fmt.Println(err)
	}
	if err2 == nil {
		cantidad += animalGato(8)
	} else {
		fmt.Println(err2)
	}
	if err3 == nil {
		cantidad += animalHamster(5)
	} else {
		fmt.Println(err3)
	}
	if err4 == nil {
		cantidad += animalTarantula(1)
	} else {
		fmt.Println(err4)
	}

	fmt.Println("Kilos de comida a comprar: ", cantidad)
}

func animal(animal string) (func(cantAnimales float64) float64, error) {
	switch animal {
	case "perro":
		return animalPerro, nil
	case "gato":
		return animalGato, nil
	case "hamster":
		return animalHamster, nil
	case "tarantula":
		return animalTarantula, nil
	default:
		return nil, errors.New("No hay " + animal + " en el refugio")
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

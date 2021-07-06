package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Ejercicio5")
	funcAnimalPerro, err := animal(PERRO)
	funcAnimalGato, err := animal(GATO)
	funcAnimalTarantula, err := animal(TARANTULA)
	funcAnimalHaster, err := animal(HAMSTER)

	var cantidadComidaComprar float64
	if err != nil {
		cantidadComidaComprar = 0
		fmt.Println("Ocurrio un error:", err)
	} else {
		cantidadComidaComprar += funcAnimalPerro(5)
		cantidadComidaComprar += funcAnimalGato(11)
		cantidadComidaComprar += funcAnimalTarantula(5)
		cantidadComidaComprar += funcAnimalHaster(6)
		fmt.Println("La cantidad de kilos de comida necesaria para todos los animales es de:", cantidadComidaComprar)
	}
}

const (
	PERRO     = "perro"
	GATO      = "gato"
	TARANTULA = "tarantula"
	HAMSTER   = "hamster"
)

func animal(animal string) (func(cantidadAnimal int) float64, error) {

	switch animal {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case TARANTULA:
		return animalTarantula, nil
	case HAMSTER:
		return animalHamster, nil
	}
	return nil, errors.New("Uno de los animales especificados todavia no se encuentra en el refugio")
}

func animalPerro(cantidadAnimal int) float64 {
	return (10 * float64(cantidadAnimal))
}
func animalGato(cantidadAnimal int) float64 {
	return 5 * float64(cantidadAnimal)
}
func animalTarantula(cantidadAnimal int) float64 {
	return 0.250 * float64(cantidadAnimal)
}
func animalHamster(cantidadAnimal int) float64 {
	return 0.150 * float64(cantidadAnimal)
}

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

func animal(tipoAn string) (func(cantAn int) float64, error) {
	switch tipoAn {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTER:
		return animalHamster, nil
	case TARANTULA:
		return animalTarantula, nil
	default:
		return nil, errors.New("tipo de animal especificado invalido")
	}
}

func animalPerro(cantAn int) float64 {
	return float64(cantAn) * 10.0
}

func animalGato(cantAn int) float64 {
	return float64(cantAn) * 5.0
}

func animalHamster(cantAn int) float64 {
	return float64(cantAn) * 0.25
}

func animalTarantula(cantAn int) float64 {
	return float64(cantAn) * 0.15
}

func main() {
	animalPerro, err := animal(PERRO)
	animalGato, err := animal(GATO)

	if err != nil {
		fmt.Println(err)
	} else {
		var cantidad float64
		cantidad += animalPerro(5)
		cantidad += animalGato(8)
		fmt.Printf("La cantidad de alimento a comprar es de %.2f KGs\n", cantidad)
	}

}

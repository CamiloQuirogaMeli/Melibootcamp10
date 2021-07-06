package main

import (
	"errors"
	"fmt"
)

const (
	TARANTULA = "tarantula"
	HAMSTER   = "hamster"
	PERRO     = "perro"
	GATO      = "gatos"
)

func calPerro(cantidad int) (result float32) {

	result = 0
	result = float32(cantidad) * 10
	return
}

func calHamster(cantidad int) (result float32) {

	result = 0
	result = float32(cantidad) * 0.25
	return
}

func calTarantula(cantidad int) (result float32) {

	result = 0
	result = float32(cantidad) * 0.15
	return
}

func calGato(cantidad int) (result float32) {

	result = 0
	result = float32(cantidad) * 5
	return
}

func animal(animalType string) (funcion func(int) float32, err error) {
	err = nil
	switch animalType {
	case PERRO:
		funcion = calPerro
	case GATO:
		funcion = calGato
	case HAMSTER:
		funcion = calHamster
	case TARANTULA:
		funcion = calTarantula
	default:
		err = errors.New("opcion incorrecta")
	}
	return
}

func main() {

	var cantidad float32
	hamFunc, hamErr := animal(HAMSTER)
	perFunc, perErr := animal(PERRO)
	gatFunc, gatErr := animal(GATO)
	tarFunc, tarErr := animal(TARANTULA)

	if hamErr == nil {
		cantidad += hamFunc(5)
	}

	if perErr == nil {
		cantidad += perFunc(5)
	}

	if gatErr == nil {
		cantidad += gatFunc(5)
	}

	if tarErr == nil {
		cantidad += tarFunc(5)
	}

	fmt.Println("Cantidad total de alimento:", cantidad)

}

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

func main() {
	var cantidad float64 = 0

	animalPerro, err := Animal(PERRO)
	if err == nil {
		cantidad += animalPerro(5)
		fmt.Println("La cantidad al momento es: ", cantidad)
	} else {
		fmt.Println(err)
	}

	animalGato, err := Animal(GATO)
	if err == nil {
		cantidad += animalGato(8)
		fmt.Println("La cantidad al momento es: ", cantidad)
	} else {
		fmt.Println(err)
	}

	animalHamster, err := Animal(HAMSTER)
	if err == nil {
		cantidad += animalHamster(5)
		fmt.Println("La cantidad al momento es: ", cantidad)
	} else {
		fmt.Println(err)
	}

	animalTarantula, err := Animal(TARANTULA)
	if err == nil {
		cantidad += animalTarantula(10)
		fmt.Println("La cantidad al momento es: ", cantidad)
	} else {
		fmt.Println(err)
	}

	animalJose, err := Animal("JOSE")
	if err == nil {
		cantidad += animalJose(10)
		fmt.Println("La cantidad al momento es: ", cantidad)
	} else {
		fmt.Println(err)
	}

}

func Animal(codigo string) (op func(cantidad int) float64, err error) {
	switch codigo {
	case PERRO:
		op = animalPerro
	case GATO:
		op = animalGato
	case HAMSTER:
		op = animalHamster
	case TARANTULA:
		op = animalTarantula
	default:
		err = errors.New("animal incorrecto")
	}
	return
}

func animalPerro(numero int) (cant float64) {
	cant = float64(numero * 10)
	return
}

func animalGato(numero int) (cant float64) {
	cant = float64(numero * 5)
	return
}

func animalHamster(numero int) (cant float64) {
	cant = float64(numero) * 0.25
	return
}

func animalTarantula(numero int) (cant float64) {
	cant = float64(numero) * 0.15
	return
}

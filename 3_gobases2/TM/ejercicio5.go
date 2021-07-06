package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalPerro, errPerr := Animal(perro)
	animalGato, errGat := Animal(gato)
	animalHamster, errHam := Animal(hamster)
	animalTarantula, errTar := Animal(tarantula)

	var cantidad float64

	if errPerr == nil {
		cantidad += animalPerro(5)
	} else {
		fmt.Println(errPerr)
	}

	if errGat == nil {
		cantidad += animalGato(8)
	} else {
		fmt.Println(errGat)
	}

	if errHam == nil {
		cantidad += animalHamster(4)
	} else {
		fmt.Println(errHam)
	}

	if errTar == nil {
		cantidad += animalTarantula(1)
	} else {
		fmt.Println(errTar)
	}

	fmt.Printf("La cantidad de alimento a comprar es de %.3f kg\n", cantidad)
}

func Animal(animal string) (func(catidad int) float64, error) {
	switch animal {
	case perro:
		return alimentoPerro, nil
	case gato:
		return alimentoGato, nil
	case hamster:
		return alimentoHamster, nil
	case tarantula:
		return alimentoTarantula, nil
	default:
		return nil, errors.New("Animal invalido")
	}
}

func alimentoPerro(nPerros int) float64 {
	return float64(nPerros) * 10
}

func alimentoGato(nGatos int) float64 {
	return float64(nGatos) * 5
}

func alimentoHamster(nHamster int) float64 {
	return float64(nHamster) * 0.250
}

func alimentoTarantula(nTarantulas int) float64 {
	return float64(nTarantulas) * 0.150
}

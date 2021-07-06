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

	var cantidad float64

	animalPerro, err := Animal(perro)
	animalGato, err1 := Animal(gato)
	animalHamster, err2 := Animal(hamster)
	animalTarantula, err3 := Animal(tarantula)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		cantidad = animalPerro(5)
		fmt.Printf("El alimento para perro son %f Kg\n ", cantidad)
	}

	if err != nil {
		fmt.Println("Error: ", err1)
	} else {
		cantidad = animalGato(8)
		fmt.Printf("El alimento para gato son %f Kg\n", cantidad)
	}

	if err != nil {
		fmt.Println("Error: ", err2)
	} else {
		cantidad = animalHamster(2)
		fmt.Printf("El alimento para hamster son %f Kg\n", cantidad)
	}

	if err != nil {
		fmt.Println("Error: ", err3)
	} else {
		cantidad = animalTarantula(7)
		fmt.Printf("El alimento para tarantula son %f Kg\n", cantidad)
	}

}

func Animal(animal string) (func(cant float64) float64, error) {
	switch animal {
	case perro:
		return cantPerr, nil
	case gato:
		return cantGat, nil
	case hamster:
		return cantHams, nil
	case tarantula:
		return cantTar, nil
	}
	return nil, errors.New("El animal no es correcto")
}

func cantPerr(can float64) float64 {
	return can * 10
}

func cantHams(can float64) float64 {
	return can * 0.250
}

func cantGat(can float64) float64 {
	return can * 5
}

func cantTar(can float64) float64 {
	return can * 0.150
}

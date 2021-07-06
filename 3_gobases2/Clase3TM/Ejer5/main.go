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
	animalPerro, err := Animal(perro)
	if err != nil {
		fmt.Println("ups se ha presentado un problema")
	} else {
		animalGato, err := Animal(gato)
		if err != nil {
			fmt.Println("ups se ha presentado un problema")
		} else {
			var cantidad float64
			cantidad += animalPerro(5)
			cantidad += animalGato(8)
			fmt.Println(cantidad, "kg")
		}
	}

}

func Animal(animal string) (func(int64) float64, error) {
	switch animal {
	case perro:
		return dog, nil
	case gato:
		return cat, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return spider, nil
	default:
		return nil, errors.New("Animal no existe")
	}
}

func dog(dogs int64) float64 {
	return 10.0 * float64(dogs)
}
func cat(cats int64) float64 {
	return 5.0 * float64(cats)
}
func spider(spiders int64) float64 {
	return 0.15 * float64(spiders)
}
func hamsterFunc(hamsters int64) float64 {
	return 0.25 * float64(hamsters)
}

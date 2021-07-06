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
	animalPerro, errPerro := Animal(perro)
	animalGato, errGato := Animal(gato)
	animalHamster, errHamster := Animal(hamster)
	animalTarantula, errTarantula := Animal(tarantula)

	fmt.Printf("Comida necesaria para 5 perros: %.2f kg\n", animalPerro(5))
	if errPerro != nil {
		fmt.Println(errPerro)
	}
	fmt.Printf("Comida necesaria para 8 gatos: %.2f kg\n", animalGato(8))
	if errGato != nil {
		fmt.Println(errGato)
	}
	fmt.Printf("Comida necesaria para 10 tarantulas: %.2f kg\n", animalTarantula(10))
	if errTarantula != nil {
		fmt.Println(errTarantula)
	}
	fmt.Printf("Comida necesaria para 20 hamsters: %.2f kg\n", animalHamster(20))
	if errHamster != nil {
		fmt.Println(errHamster)
	}
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case perro:
		return cantidadPerro, nil
	case gato:
		return cantidadGato, nil
	case hamster:
		return cantidadHamster, nil
	case tarantula:
		return cantidadTarantula, nil
	}
	return nil, errors.New("error: animal no especificado")
}

func cantidadPerro(perros int) float64 {
	return float64(perros) * 10
}

func cantidadGato(gatos int) float64 {
	return float64(gatos) * 5
}

func cantidadHamster(hamsters int) float64 {
	return float64(hamsters) * 0.25
}

func cantidadTarantula(tarantulas int) float64 {
	return float64(tarantulas) * 0.15
}

package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	TARANTULA = "tarantula"
)

func main() {
	perro, error := animal(PERRO)
	gato, error := animal(GATO)
	tarantula, error := animal(TARANTULA)

	if error != nil {
		fmt.Println("Se produjo un error")
	} else {
		valorPerro := perro(5)
		valorGato := gato(3)
		valorTarantula := tarantula(6)

		fmt.Println("valorMinimo ", valorPerro)
		fmt.Println("valorMaximo ", valorGato)
		fmt.Println("valorPromedio ", valorTarantula)
	}

}

func animal(op string) (func(float64) float64, error) {
	switch op {
	case PERRO:
		return perro, nil
	case GATO:
		return gato, nil
	case TARANTULA:
		return tarantula, nil
	default:
		return nil, errors.New("No existe la función")
	}
}

func perro(cant float64) float64 {
	return 10 * cant
}

func gato(cant float64) float64 {
	return 5 * cant
}

func tarantula(cant float64) float64 {
	return 0.15 * cant
}

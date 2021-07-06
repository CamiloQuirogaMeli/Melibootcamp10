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
	animal, err := Animal(TARANTULA)
	fmt.Println("Error", err)
	cantidad := animal(1)
	fmt.Println("Cantidad", cantidad)
}

func Animal(tipo string) (func(cantidad int) float64, error) {
	switch tipo {
	case PERRO:
		return perro, nil
	case GATO:
		return gato, nil
	case HAMSTER:
		return hamster, nil
	case TARANTULA:
		return tarantula, nil
	default:
		return nil, errors.New("No existe el animal especificado")
	}
}

func perro(cantidad int) float64 {
	return float64(cantidad) * 10.0
}

func gato(cantidad int) float64 {
	return float64(cantidad) * 5.0
}

func hamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func tarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)

	genFunc, err := generarOperacion(minimo)

	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	valor, err := genFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	fmt.Printf("El valor obtenido para la operacion seleccionada es: %f \n", valor)

}

func generarOperacion(operacion string) (func(valores ...int) (float32, error), error) {

	switch operacion {
	case "minimo":
		return obtenerMin, nil
	case "maximo":
		return obtenerMaximo, nil
	case "promedio":
		return calcularPromedio, nil
	default:
		return nil, errors.New("No esta definida la operación")

	}
}

func obtenerMaximo(notas ...int) (float32, error) {
	max := 0

	for _, val := range notas {
		if val < 0 {
			return -1, errors.New("Se encontro una nota negativa ")
		}
		if val > max {
			max = val
		}
	}

	return float32(max), nil
}

func obtenerMin(notas ...int) (float32, error) {
	min := notas[0]

	for _, val := range notas {
		if val < 0 {
			return -1, errors.New("Se encontro una nota negativa ")
		}
		if val < min {
			min = val
		}
	}

	return float32(min), nil
}

func calcularPromedio(notas ...int) (float32, error) {
	count := 0

	for _, val := range notas {
		if val < 0 {
			return -1, errors.New("Se encontro una nota negativa ")
		}

		count += val
	}

	return float32(count) / float32(len(notas)), nil

}

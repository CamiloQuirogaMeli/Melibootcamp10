package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func main() {
	minFunc, err := operacion(MINIMO)
	fmt.Printf("MinFunc: %T err: %v\n", minFunc, err)
	maxFunc, err := operacion(MAXIMO)
	fmt.Printf("MaxFunc: %T err: %v\n", maxFunc, err)
	promFunc, err := operacion(PROMEDIO)
	fmt.Printf("PromFunc: %T err: %v\n", promFunc, err)
	errorFunc, err := operacion("ERROR")
	fmt.Printf("ErrorFunc: %T err: %v\n", errorFunc, err)
	valorMinimo := minFunc(1, 2, 3, 4, 5)
	fmt.Println("ValorMinimo:", valorMinimo)
	valorMaximo := maxFunc(1, 2, 3, 4, 5)
	fmt.Println("ValorMaximo:", valorMaximo)
	valorPromedio := promFunc(1, 2, 3, 4, 5)
	fmt.Println("ValorPromedio:", valorPromedio)
}

func operacion(operacion string) (func(datos ...int) float64, error) {
	switch operacion {
	case MINIMO:
		return minimo, nil
	case MAXIMO:
		return maximo, nil
	case PROMEDIO:
		return promedio, nil
	default:
		return nil, errors.New("No existe la operación")
	}
}

func promedio(datos ...int) float64 {
	var p float64
	for _, valor := range datos {
		p += float64(valor)
	}
	return p / float64(len(datos))
}

func maximo(datos ...int) float64 {
	var m float64
	for indice, valor := range datos {
		if indice == 0 || float64(valor) > m {
			m = float64(valor)
		}
	}
	return m
}

func minimo(datos ...int) float64 {
	var m float64
	for indice, valor := range datos {
		if indice == 0 || float64(valor) < m {
			m = float64(valor)
		}
	}
	return m
}

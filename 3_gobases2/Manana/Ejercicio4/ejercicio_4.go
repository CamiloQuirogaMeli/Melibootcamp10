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

func operacion(tipoOp string) (func(valores ...int) float64, error) {
	switch tipoOp {
	case MINIMO:
		return minFunc, nil
	case PROMEDIO:
		return promFunc, nil
	case MAXIMO:
		return maxFunc, nil
	default:
		return nil, errors.New("tipo de operacion especificada invalida")
	}
}

func minFunc(valores ...int) float64 {
	var min int
	for i, valor := range valores {
		if i == 0 {
			min = valor
		} else {
			if valor < min {
				min = valor
			}
		}
	}
	return float64(min)
}

func maxFunc(valores ...int) float64 {
	var max int
	for i, valor := range valores {
		if i == 0 {
			max = valor
		} else {
			if valor > max {
				max = valor
			}
		}
	}
	return float64(max)
}

func promFunc(valores ...int) float64 {
	sum := 0
	for _, valor := range valores {
		sum += valor
	}
	return float64(sum) / float64(len(valores))
}

func main() {
	minFunc, err := operacion(MINIMO)
	promFunc, err := operacion(PROMEDIO)
	maxFunc, err := operacion(MAXIMO)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valor minimo: ", valorMinimo)
		fmt.Println("Valor maximo: ", valorMaximo)
		fmt.Printf("Valor promedio: %.2f\n", valorPromedio)
	}
}

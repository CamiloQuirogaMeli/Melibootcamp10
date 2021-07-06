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

func calcularEstadisticas(calculo string) (func(valores ...int) int, error) {
	switch calculo {
	case MINIMO:
		return calcularMinimo, nil
	case PROMEDIO:
		return calcularPromedio, nil
	case MAXIMO:
		return calcularMaximo, nil
	default:
		return nil, errors.New("Selecciona un tipo de calculo valido")
	}
}

func calcularMinimo(valores ...int) int {
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
	return min
}

func calcularMaximo(valores ...int) int {
	var max int
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}

	return max
}

func calcularPromedio(valores ...int) int {
	var promedio int
	for _, valor := range valores {
		promedio += valor
	}
	promedio /= len(valores)

	return promedio
}

func main() {
	minFunc, err := calcularEstadisticas(MINIMO)
	promFunc, err := calcularEstadisticas(PROMEDIO)
	maxFunc, err := calcularEstadisticas(MAXIMO)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		//Si hubo error
		fmt.Println(err)
	} else {
		fmt.Println("El minimo es:", valorMinimo)
		fmt.Println("El promedio es:", valorPromedio)
		fmt.Println("El maximo es:", valorMaximo)
	}
}

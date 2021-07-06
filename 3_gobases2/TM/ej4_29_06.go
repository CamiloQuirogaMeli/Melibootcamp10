package main

import (
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func operacion(operador string) func(valores ...float64) float64 {
	switch operador {
	case minimo:
		return minFunc
	case promedio:
		return promFunc
	case maximo:
		return maxFunc
	}
	return nil
}

func minFunc(valores ...float64) float64 {
	var min float64 = valores[0]
	for _, value := range valores {
		if value < min {
			min = value
		}
	}
	return min
}

func maxFunc(valores ...float64) float64 {
	var max float64 = 0
	for _, value := range valores {
		if value > max {
			max = value
		}
	}
	return max
}

func promFunc(valores ...float64) float64 {
	var resultado float64

	for _, valor := range valores {
		if valor < 0 {
			return 0
		} else {
			resultado += valor
		}
	}
	return resultado / float64(len(valores))
}

func main() {

	minFunc := operacion(minimo)
	promFunc := operacion(promedio)
	maxFunc := operacion(maximo)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("maximo: ", valorMaximo)
	fmt.Println("promedio: ", valorPromedio)
	fmt.Println("minimo: ", valorMinimo)

}

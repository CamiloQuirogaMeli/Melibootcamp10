package main

import (
	"fmt"
)

const (
	MINIMO   = "MINIMO"
	PROMEDIO = "PROMEDIO"
	MAXIMO   = "MAXIMO"
)

func main() {
	minFunc := operacion(MINIMO)
	promFunc := operacion(PROMEDIO)
	maxFunc := operacion(MAXIMO)

	//minFunc, err := operacion(MINIMO)
	//promFunc, err := operacion(PROMEDIO)
	//maxFunc, err := operacion(MAXIMO)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Valor minimo: ", valorMinimo)
	fmt.Println("Valor maximo: ", valorMaximo)
	fmt.Printf("Valor promedio: %.2f\n", valorPromedio)

}

func operacion(operador string) func(values ...float64) float64 {
	switch operador {
	case MINIMO:
		return minFunc
	case PROMEDIO:
		return promFunc
	case MAXIMO:
		return maxFunc
	default:
		return nil
	}

}

func minFunc(values ...float64) float64 {
	var min float64 = values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maxFunc(values ...float64) float64 {
	var max float64 = values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func promFunc(values ...float64) float64 {
	var prom float64 = values[0]
	for _, value := range values {
		prom += value
	}
	return prom / float64(len(values))
}

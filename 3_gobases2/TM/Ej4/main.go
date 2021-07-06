package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	minFunc, err1 := operacion(minimo)
	promFunc, err2 := operacion(promedio)
	maxFunc, err3 := operacion(maximo)

	valorMinimo := minFunc(2, 3, -3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, -4, 5)

	fmt.Println("Valor minimo: ", valorMinimo)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Promedio promedio: ", valorPromedio)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("Valor maximo: ", valorMaximo)

	if err3 != nil {
		fmt.Println(err3)
	}
}

func operacion(op string) (func(values ...float64) float64, error) {
	switch op {
	case minimo:
		return opMin, nil
	case promedio:
		return opProm, nil
	case maximo:
		return opMax, nil
	default:
		return nil, errors.New("el cálculo no está definido")
	}
}

func opMin(values ...float64) float64 {
	var min float64
	if len(values) > 0 {
		min = values[0]
	} else {
		return 0
	}
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func opMax(values ...float64) float64 {
	var max float64
	if len(values) > 0 {
		max = values[0]
	} else {
		return 0
	}
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func opProm(values ...float64) float64 {
	var prom float64
	for _, value := range values {
		prom += value
	}
	if len(values) > 0 {
		prom /= float64(len(values))
	}
	return prom
}

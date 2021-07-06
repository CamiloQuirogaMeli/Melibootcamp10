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
	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Println("operacion no existe")
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println(valorMinimo)
		fmt.Println(valorPromedio)
		fmt.Println(valorMaximo)
	}
}

func operacion(operation string) (func(value ...float64) float64, error) {
	switch operation {
	case minimo:
		return minimunValue, nil
	case promedio:
		return averageValue, nil
	case maximo:
		return maximunValue, nil
	default:
		return nil, errors.New("operacion no soportada")
	}
}

func maximunValue(values ...float64) float64 {
	var max float64
	for i, value := range values {
		if i == 0 {
			max = value
		} else {
			if max < value {
				max = value
			}
		}
	}
	return max
}
func minimunValue(values ...float64) float64 {
	var min float64
	for i, value := range values {
		if i == 0 {
			min = value
		} else {
			if min > value {
				min = value
			}
		}
	}
	return min
}
func averageValue(values ...float64) float64 {
	avg := 0.0
	for _, value := range values {
		avg += value
	}
	return avg / float64(len(values))
}

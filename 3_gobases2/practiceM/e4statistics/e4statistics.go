package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minFuncimo"
	PROMEDIO = "promFuncedio"
	MAXIMO   = "maxfuncimo"
)

func main() {
	minFunc, error := operacion(MINIMO)
	maxFunc, error := operacion(MAXIMO)
	promFunc, error := operacion(PROMEDIO)

	if error != nil {
		fmt.Println("Se produjo un error")
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("valorMinimo ", valorMinimo)
		fmt.Println("valorMaximo ", valorMaximo)
		fmt.Println("valorPromedio ", valorPromedio)
	}

}

func operacion(op string) (func(...int) float64, error) {
	switch op {
	case MINIMO:
		return minFunc, nil
	case MAXIMO:
		return maxfunc, nil
	case PROMEDIO:
		return promFunc, nil
	default:
		return nil, errors.New("No existe la función")
	}
}

func maxfunc(values ...int) float64 {
	var max int
	flag := true
	for _, value := range values {
		if flag {
			max = value
			flag = false
		} else {
			if value > max {
				max = value
			}
		}
	}
	return (float64(max))
}

func minFunc(values ...int) float64 {
	var min int
	flag := true
	for _, value := range values {
		if flag {
			min = value
			flag = false
		} else {
			if value < min {
				min = value
			}
		}
	}
	return float64(min)
}

func promFunc(values ...int) float64 {
	var count int
	for _, value := range values {
		count = count + value
	}
	return float64(count / len(values))
}

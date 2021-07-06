package main

import (
	"errors"
	"fmt"
)

func main() {

	const (
		MINIMO   = "minimo"
		PROMEDIO = "promedio"
		MAXIMO   = "maximo"
	)

	minFunc, err := operation(MINIMO)
	promFunc, err2 := operation(PROMEDIO)
	maxFunc, err3 := operation(MAXIMO)

	if err == nil {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor minimo: ", valorMinimo)
	} else {
		fmt.Println(err)
	}
	if err2 == nil {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor promedio: ", valorPromedio)
	} else {
		fmt.Println(err2)
	}
	if err3 == nil {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor maximo: ", valorMaximo)
	} else {
		fmt.Println(err3)
	}
}

func operation(operation string) (func(qualifications ...int) int, error) {

	switch operation {
	case "minimo":
		return minFunc, nil
	case "promedio":
		return promFunc, nil
	case "maximo":
		return maxFunc, nil
	default:
		return nil, errors.New("Calculo " + operation + " no definido")
	}
}

func minFunc(qualifications ...int) int {
	min := qualifications[0]
	for _, qualification := range qualifications {
		if qualification < min {
			min = qualification
		}
	}

	return min
}

func promFunc(qualifications ...int) int {
	var total int
	for _, qualification := range qualifications {
		total += qualification
	}

	return total / len(qualifications)
}

func maxFunc(qualifications ...int) int {
	max := qualifications[0]
	for _, qualification := range qualifications {
		if qualification > max {
			max = qualification
		}
	}

	return max
}

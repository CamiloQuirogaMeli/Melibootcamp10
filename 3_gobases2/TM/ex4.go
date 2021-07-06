package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimo"
	maximum = "maximo"
	average = "promedio"
)

func getMaximum(values ...float64) (max float64) {
	for i, v := range values {
		if i == 0 || v > max {
			max = v
		}
	}
	return
}

func getMinimum(values ...float64) (min float64) {
	for i, v := range values {
		if i == 0 || v < min {
			min = v
		}
	}
	return
}

func getAverage(values ...float64) (avg float64) {
	for _, v := range values {
		avg += v
	}
	avg = avg / float64(len(values))
	return
}

func generate_opertation_function(operation string) (func(values ...float64) float64, error) {
	switch operation {
	case minimum:
		return getMinimum, nil
	case maximum:
		return getMaximum, nil
	case average:
		return getAverage, nil
	default:
		return nil, errors.New("Operación incorrecta")
	}
}

func main() {
	data := []float64{1, 2, 4, 5, 11, 2, 5, 6, 7}
	operation := maximum
	function, err := generate_opertation_function(operation)
	if err == nil {
		fmt.Println("Se llamo la operación", operation, "con los datos ...")
		fmt.Println("Y el resultado fue:", function(data...))
	} else {
		fmt.Println(err)
	}

	//Tests described in the excercise statement

	maxFunc, err := generate_opertation_function(maximum)
	minFunc, err := generate_opertation_function(minimum)
	avgFunc, err := generate_opertation_function(average)
	minimunValue := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	avergaeValue := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maximumValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minimunValue, maximumValue, avergaeValue)
}

package main

import (
	"errors"
	"fmt"
)

const (
	MIN = "min"
	AVG = "avg"
	MAX = "max"
)

func main() {
	minFunc, err := operation(MIN)
	avgFunc, err := operation(AVG)
	maxFunc, err := operation(MAX)

	if err == nil {
		minValue := minFunc(2, 5, 8, 6, 5, 4, 67)
		avgValue := avgFunc(5, 10)
		maxValue := maxFunc(2, 5, 8, 6, 5, 4, 67)

		fmt.Println("La menor nota es: ", minValue)

		fmt.Println("El promedio es: ", avgValue)

		fmt.Println("La mayor nota es: ", maxValue)
	} else {
		fmt.Println(err)
	}
}

func operation(operation string) (func(values ...int) float64, error) {

	switch operation {
	case MIN:
		return getMin, nil
	case AVG:
		return getAvg, nil
	case MAX:
		return getMax, nil
	}
	return nil, errors.New("ERROR! El calculo no está definido")
}

func getMin(values ...int) float64 {
	var min float64
	min = float64(values[0])
	for _, element := range values {
		if float64(element) < min {
			min = float64(element)
		}
	}
	return min
}

func getAvg(values ...int) float64 {
	var avg float64
	var sum, count int = 0, 0

	for _, element := range values {
		sum += element
		count++
	}
	avg = float64(sum) / float64(count)
	return avg
}

func getMax(values ...int) float64 {
	var max float64
	max = float64(values[0])
	for _, element := range values {
		if float64(element) > max {
			max = float64(element)
		}
	}
	return max
}

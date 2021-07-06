package main

import (
	"errors"
	"fmt"
)

const (
	MIN = "minimo"
	MAX = "maximo"
	AVG = "promedio"
)

func main() {
	min, max, avg := "minimo", "maximo", "promedio"

	minFunc, _ := getOperation(min)
	maxFunc, _ := getOperation(max)
	avgFunc, _ := getOperation(avg)

	anotherOperation, errOperation := getOperation("medium")

	minValue, _ := minFunc(10, 20, 30, 50, 5, 2, 100)
	maxValue, _ := maxFunc(10, 20, 30, 50, 5, 2, 100)
	avgValue, errAvg := avgFunc(10, 20, 30, 50, 5, 2, 100)

	fmt.Printf("Min [%0.f]\n", minValue)
	fmt.Printf("Max [%0.f]\n", maxValue)
	fmt.Printf("Avg [%2.f]\n error [%v]", avgValue, errAvg)

	if anotherOperation == nil {
		fmt.Println(errOperation)
	}

}

func getOperation(operation string) (func(califications ...int) (float64, error), error) {
	switch operation {
	case MIN:
		return findMin, nil
	case MAX:
		return findMax, nil
	case AVG:
		return avg, nil
	}
	return nil, errors.New(fmt.Sprintf("Not found operation: [%s]\n", operation))
}

func findMin(califications ...int) (float64, error) {
	min := califications[0]

	for i := 1; i < len(califications); i++ {
		if califications[i] < min {
			min = califications[i]
		}
	}
	return float64(min), nil
}

func findMax(califications ...int) (float64, error) {
	max := califications[0]

	for i := 1; i < len(califications); i++ {
		if califications[i] > max {
			max = califications[i]
		}
	}
	return float64(max), nil
}

func avg(califications ...int) (float64, error) {
	var avg, sum float64
	for _, calification := range califications {
		if calification < 0 {
			return 0, errors.New("A calification is negative.")
		}
		sum += float64(calification)
	}
	avg = divide(sum, float64(len(califications)))
	return avg, nil
}

func divide(sumCalifications float64, countCalifications float64) float64 {
	return sumCalifications / countCalifications
}

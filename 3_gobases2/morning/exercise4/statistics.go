package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	min  = "min"
	mean = "mean"
	max  = "max"
)

func main() {

	print := fmt.Println

	operation := "mean"
	operate, err := operationHandler(operation)
	if err != nil {
		print(err)
	} else {
		result := operate(5, 4, 2, 5.4)
		print(result)
	}
}

func getMin(nums ...float64) float64 {
	currentMin := math.Inf(1)

	for _, num := range nums {
		currentMin = math.Min(num, currentMin)
	}
	return currentMin
}

func getMax(nums ...float64) float64 {
	currentMax := math.Inf(-1)

	for _, num := range nums {
		currentMax = math.Max(num, currentMax)
	}
	return currentMax
}

func getMean(nums ...float64) float64 {

	mean := 0.0

	for _, grade := range nums {
		mean += grade
	}
	return mean / float64(len(nums))
}

func operationHandler(operation string) (func(nums ...float64) float64, error) {

	switch operation {
	case min:
		return getMin, nil
	case max:
		return getMax, nil
	case mean:
		return getMean, nil
	default:
		return nil, errors.New("function " + operation + "is not defined")
	}
}

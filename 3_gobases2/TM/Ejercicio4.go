package main

import (
	"errors"
	"fmt"
)

const (
	MIN = "min"
	MAX = "max"
	AVG = "avg"
)

func main() {
	function, err := operations(MIN)
	if err != nil {
		fmt.Println(err)
	} else {
		result := function(1, 10)
		fmt.Println(result)
	}
}

func operations(op string) (func(val ...int) float64, error) {
	switch op {
	case MIN:
		return minFunction, nil
	case MAX:
		return maxFunction, nil
	case AVG:
		return avgFunction, nil
	default:
		return nil, errors.New("El calculo no esta definido")
	}
}

func minFunction(values ...int) float64 {
	minValue := 999999
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}
	return float64(minValue)
}

func maxFunction(values ...int) float64 {
	maxValue := -999999
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return float64(maxValue)
}

func avgFunction(notes ...int) float64 {
	var total float64 = 0
	for _, value := range notes {
		total += float64(value)
	}
	return total / float64(len(notes))
}

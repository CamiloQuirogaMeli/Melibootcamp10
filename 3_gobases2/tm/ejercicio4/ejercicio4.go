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

	minFunc, err := operation(minimo)
	maxFunc, err := operation(maximo)
	avgFunc, err := operation(promedio)

	if err != nil {
		fmt.Println(err)
	} else {
		valorMin := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(valorMin)
		valorMax := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(valorMax)
		valorAvg := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(valorAvg)
	}
}

func operation(op string) (func(values ...int) float32, error) {
	switch op {
	case minimo:
		return minFunc, nil
	case maximo:
		return maxFunc, nil
	case promedio:
		return avgFunc, nil
	}
	return nil, errors.New("Error")
}

func minFunc(values ...int) float32 {
	min := 100
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return float32(min)
}

func maxFunc(values ...int) float32 {
	max := 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return float32(max)
}

func avgFunc(values ...int) float32 {
	sum := float32(0)
	len := float32(0)
	for _, value := range values {
		sum += float32(value)
		len++
	}
	return sum / len
}

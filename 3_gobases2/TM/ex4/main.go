package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func main() {
	minFunc, err := operacion("minimo")
	promFunc, err := operacion("promedio")
	maxFunc, err := operacion("maximo")

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err == nil {
		fmt.Println("Valor minimo: ", valorMinimo)
		fmt.Println("Valor promedio: ", valorPromedio)
		fmt.Println("Valor maximo: ", valorMaximo)
	}
}

func operacion(inst string) (func(nums ...float64) float64, error) {
	switch inst {
	case MINIMO:
		return minOper, nil
	case PROMEDIO:
		return promOper, nil
	case MAXIMO:
		return maxOper, nil
	}
	return nil, errors.New("inst invalida")
}

func minOper(nums ...float64) float64 {
	var rmin float64 = nums[0]
	for _, num := range nums {
		if num < rmin {
			rmin = num
		}
	}
	return rmin
}

func promOper(nums ...float64) float64 {
	var size int = 0
	var sum float64 = 0
	for _, num := range nums {
		if num < 0 {
			return float64(-1)
		}
		sum += num
		size++
	}
	return float64(sum) / float64(size)
}

func maxOper(nums ...float64) float64 {
	var rmax float64 = nums[0]
	for _, num := range nums {
		if num > rmax {
			rmax = num
		}
	}
	return rmax
}

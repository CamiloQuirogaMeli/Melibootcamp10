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
	minFunc, errMin := operacion(minimo)
	promFunc, errProm := operacion(promedio)
	maxFunc, errMax := operacion(maximo)

	if errMin != nil {
		fmt.Println(errMin)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Minimo:", valorMinimo)
	}
	if errMax != nil {
		fmt.Println(errMax)
	} else {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Maximo:", valorMaximo)
	}
	if errProm != nil {
		fmt.Println(errProm)
	} else {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Promedio:", valorPromedio)
	}
}

func operacion(op string) (func(values ...int) float64, error) {
	switch op {
	case minimo:
		return min, nil
	case maximo:
		return max, nil
	case promedio:
		return prom, nil
	default:
		return nil, errors.New("operacion no valida")
	}
}

func min(values ...int) float64 {
	var minVal float64 = float64(values[0])
	for _, val := range values {
		if float64(val) < minVal {
			minVal = float64(val)
		}
	}
	return minVal
}

func max(values ...int) float64 {
	var maxVal float64 = float64(values[0])
	for _, val := range values {
		if float64(val) > maxVal {
			maxVal = float64(val)
		}
	}
	return maxVal
}

func prom(ratings ...int) float64 {
	var ratingsSum float64
	for _, rating := range ratings {
		ratingsSum += float64(rating)
	}
	return ratingsSum / float64(len(ratings))
}

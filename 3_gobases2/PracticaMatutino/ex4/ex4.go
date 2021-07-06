package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	min = "minimum"
	max = "maximum"
	avg = "average"
)

func main() {

	minFnc := getFunction(min)

	minResult, minErr := minFnc(6.0, 5.3, 6.7, 2.3, 4.1, 7.0, 2.1)

	if minErr == nil {
		fmt.Println("Min result:")
		fmt.Println(minResult)
	} else {
		fmt.Println(minErr)
	}

	maxFnc := getFunction(max)

	maxResult, maxErr := maxFnc(6.0, 5.3, 6.7, 2.3, 4.1, 7.0, 2.1)

	if maxErr == nil {
		fmt.Println("Max result:")
		fmt.Println(maxResult)
	} else {
		fmt.Println(maxErr)
	}

	avgFnc := getFunction(avg)

	avgResult, avgErr := avgFnc(6.0, 5.3, 6.7, 2.3, 4.1, 7.0, 2.1)

	if avgErr == nil {
		fmt.Println("Avg result:")
		fmt.Println(avgResult)
	} else {
		fmt.Println(avgErr)
	}

}

func getFunction(opc string) func(values ...float64) (float64, error) {

	switch opc {
	case "minimum":
		return getMin
	case "maximum":
		return getMax
	case "average":
		return getAvg
	default:
		return nil
	}
}

func getMin(values ...float64) (float64, error) {

	if isEmpty(values) {
		return 0.0, errors.New("empty values")
	}

	minValue := math.MaxFloat64

	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue, nil
}

func getMax(values ...float64) (float64, error) {

	if isEmpty(values) {
		return 0.0, errors.New("empty values")
	}

	maxValue := 0.0

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue, nil
}

func getAvg(values ...float64) (float64, error) {

	if isEmpty(values) {
		return 0.0, errors.New("empty values")
	}

	sum := 0.0

	for _, value := range values {
		sum += value
	}

	return (sum / (float64)(len(values))), nil
}

func isEmpty(aSlice []float64) bool {
	return len(aSlice) == 0
}

/*
 Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
requiriendo calcular el mínimo, máximo y promedio.

 Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior
*/

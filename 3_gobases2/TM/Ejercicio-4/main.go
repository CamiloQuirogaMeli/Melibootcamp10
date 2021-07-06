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

func calcMin(values ...int) float64 {
	var minValue int = -1
	for _, value := range values {
		if minValue == -1 {
			minValue = value
		}
		if value < minValue {
			minValue = value
		}
	}
	return float64(minValue)
}

func calcAvg(values ...int) float64 {
	acumValues := 0
	for _, value := range values {
		acumValues += value
	}
	return float64(acumValues) / float64(len(values))
}

func calcMax(values ...int) float64 {
	maxValue := 0
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return float64(maxValue)
}

func operacion(typeOp string) (func(values ...int) float64, error) {

	switch typeOp {
	case "minimo":
		{
			return calcMin, nil
		}
	case "promedio":
		{
			return calcAvg, nil
		}
	case "maximo":
		{
			return calcMax, nil
		}
	default:
		return nil, errors.New("no existe la operacion solicitada")
	}
}

func main() {

	minFunc, err := operacion(minimo)
	if err != nil {
		fmt.Println(err.Error())

	}
	promFunc, err := operacion(promedio)
	if err != nil {
		fmt.Println(err.Error())

	}
	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Println(err.Error())

	}

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Valor minimo: %.0f \n Valor Promedio: %.2f \n Valor Maximo: %.0f", valorMinimo, valorPromedio, valorMaximo)

}

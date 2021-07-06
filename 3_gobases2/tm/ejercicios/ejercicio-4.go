package ejercicios

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func calcularMinimo(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func calcularMaximo(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func calcularPromedio(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum / len(values)
}

func calcular(tipoCalculo string) (func(value ...int) int, error) {
	switch tipoCalculo {
	case maximo:
		return calcularMaximo, nil

	case minimo:
		return calcularMinimo, nil

	case promedio:
		return calcularPromedio, nil

	default:
		return nil, errors.New("tipo de calculo no válido")
	}
}

func handleErrors(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CuartoEjercicio() {
	minFunc, err := calcular(minimo)
	handleErrors(err)

	promFunc, err := calcular(promedio)
	handleErrors(err)

	maxFunc, err := calcular(maximo)
	handleErrors(err)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("%s => %d\n", minimo, valorMinimo)
	fmt.Printf("%s => %d\n", promedio, valorPromedio)
	fmt.Printf("%s => %d\n", maximo, valorMaximo)
}

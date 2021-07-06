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

	minFunc, err := operacion(MINIMO)
	if err == nil {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("El valor mínimo es: ", valorMinimo)
	} else {
		fmt.Println(err)
	}

	promFunc, err := operacion(PROMEDIO)
	if err == nil {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("El promedio es: ", valorPromedio)
	} else {
		fmt.Println(err)
	}

	maxFunc, err := operacion(MAXIMO)
	if err == nil {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("El valor máximo es: ", valorMaximo)
	} else {
		fmt.Println(err)
	}

}

func operacion(codigo string) (op func(valores ...int) float64,
	err error) {
	switch codigo {
	case MINIMO:
		op = minFunc
	case PROMEDIO:
		op = promFunc
	case MAXIMO:
		op = maxFunc
	default:
		err = errors.New("operacion incorrecta")
	}
	return
}

func minFunc(numeros ...int) (minimo float64) {
	var min float64 = 150
	for _, num := range numeros {
		if float64(num) < min {
			min = float64(num)
		}
	}
	minimo = min

	return
}

func promFunc(numeros ...int) (promedio float64) {
	var total int = 0
	for _, num := range numeros {
		total += num
	}
	promedio = float64(total) / float64(len(numeros))
	return
}

func maxFunc(numeros ...int) (maximo float64) {
	var max float64 = -1
	for _, num := range numeros {
		if float64(num) > max {
			max = float64(num)
		}
	}
	maximo = max
	return
}

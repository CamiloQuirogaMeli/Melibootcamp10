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

func minFunc(valores ...int) int {
	minimo := valores[0]
	for _, valor := range valores {
		if valor < minimo {
			minimo = valor
		}
	}
	return minimo
}

func maxFunc(valores ...int) int {
	maximo := valores[0]
	for _, valor := range valores {
		if valor > maximo {
			maximo = valor
		}
	}
	return maximo
}

func promFunc(valores ...int) int {
	var resultado int
	for _, value := range valores {
		if value < 0 {
			continue
		}
		resultado += value
	}
	return resultado / len(valores)
}

func operacion(tipoDeCalculo string) (func(valores ...int) int, error) {
	switch tipoDeCalculo {
	case "maximo":
		return maxFunc, nil
	case "minimo":
		return minFunc, nil
	case "promedio":
		return promFunc, nil
	default:
		return nil, errors.New("ingrese una operacion valida")
	}
}

func main() {
	minFunc, err := operacion(minimo)
	if err != nil {
		fmt.Println(err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Minimo", valorMinimo)
	}

	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Println(err)
	} else {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Maximo", valorMaximo)
	}

	promFunc, err := operacion(promedio)
	if err != nil {
		fmt.Println(err)
	} else {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor Promedio", valorPromedio)
	}
}

package main

import (
	"errors"
	"fmt"
)

//funcion que indique que tipo de calculo queremos realizar(minimo, maximo, promedio), y nos devuelva otra funcion(error en caso en que no este definido);
//ademas, pasar cantidad N de enteros
//devolver lo solicitado

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func operacion(operacion string) (func(notas ...float32) float32, error) {
	switch operacion {
	case "minimo":
		return minFunc, nil
	case "promedio":
		return promFunc, nil
	case "maximo":
		return maxFunc, nil
	default:
		return nil, errors.New("no se encuentra la operacion")
	}
}

func minFunc(valores ...float32) float32 {
	var min float32 = valores[0]

	for _, val := range valores {
		if val < min {
			min = val
		}
	}
	return min
}

func promFunc(valores ...float32) float32 {
	var suma float32

	for _, val := range valores {
		suma += val
	}
	return suma / float32(len(valores))
}

func maxFunc(valores ...float32) float32 {
	var max float32 = valores[0]

	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {

	minFunc, err := operacion(MINIMO)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		valorMinimo := minFunc(2, 12)
		fmt.Printf("El valor minimo es: %.2f\n", valorMinimo)
	}

	promFunc, err := operacion(PROMEDIO)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		valorPromedio := promFunc(2, 12)
		fmt.Printf("El valor promedio es: %.2f\n", valorPromedio)
	}

	maxFunc, err := operacion(MAXIMO)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		valorMaximo := maxFunc(2, 12)
		fmt.Printf("El valor maximo es: %.2f\n", valorMaximo)
	}
}

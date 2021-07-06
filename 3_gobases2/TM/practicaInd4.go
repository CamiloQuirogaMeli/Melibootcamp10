package main

import (
	"errors"
	"fmt"
)

/* Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior */

func getMin(values ...int) int {

	result := 0

	for i, v := range values {
		if i != 0 {
			if v < result {
				result = v
			}
		} else {
			result = v
		}
	}

	return result

}

func getMax(values ...int) int {
	result := 0

	for _, v := range values {
		if v > result {
			result = v
		}
	}
	return result

}

func getProm(values ...int) int {
	result := 0
	j := 0

	for _, v := range values {

		result += v
		j++
	}
	result = result / j
	return result

}

func operacion(calc string) (func(values ...int) int, error) {

	switch calc {

	case MIN:
		return getMin, nil

	case MAX:
		return getMax, nil

	case PROM:
		return getProm, nil

	default:
		return nil, errors.New("El valor ingresado es incorrecto")
	}
}

const (
	MIN  = "minimo"
	MAX  = "maximo"
	PROM = "promedio"
)

func main() {

	min, err := operacion(MIN)
	max, err := operacion(MAX)
	prom, err := operacion(PROM)

	resultadoMin := min(2, 1, 3, 5, 6, 5, 6, 7)
	resultadoMax := max(2, 1, 3, 5, 6, 5, 6, 7)
	resultadoProm := prom(2, 1, 3, 5, 6, 5, 6, 7)

	if err == nil {
		fmt.Printf("El resultado de las operaciones son: \n %v es el mínimo \n %v es el máximo \n %v es el promedio", resultadoMin, resultadoMax, resultadoProm)
	}

}

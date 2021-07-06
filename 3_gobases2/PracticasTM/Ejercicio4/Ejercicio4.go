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
	/*
		Ejercicio 4 - Calcular estadísticas
		Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
		requiriendo calcular el mínimo, máximo y promedio.
		Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
		máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
		definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
		indicamos en la función anterior
	*/

	min, err := operacion(MINIMO)
	max, err := operacion(MAXIMO)
	pro, err := operacion(PROMEDIO)

	if err != nil {
		fmt.Println(err)
	} else {
		valorMinimo := min(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := max(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := pro(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println(
			"Valor Minimo:", valorMinimo, "\n",
			"Valor Promedio:", valorPromedio, "\n",
			"Valor Maximo:", valorMaximo)
	}

}

func operacion(op string) (func(...int) float64, error) {
	switch op {
	case MINIMO:
		return minFunc, nil

	case MAXIMO:
		return maxFunc, nil

	case PROMEDIO:
		return promFunc, nil
	default:
		return nil, errors.New("Calculo no definido")
	}
}

func minFunc(values ...int) float64 {
	var min int
	aux := true
	for _, value := range values {
		if aux {
			min = value
			aux = false
		} else {
			if value < min {
				min = value
			}
		}
	}
	return float64(min)
}

func maxFunc(values ...int) float64 {
	var max int
	aux := true
	for _, value := range values {
		if aux {
			max = value
			aux = false
		} else {
			if value > max {
				max = value
			}
		}
	}
	return float64(max)
}

func promFunc(values ...int) float64 {
	var sum int
	for _, value := range values {
		sum += value
	}
	return float64(sum / len(values))
}

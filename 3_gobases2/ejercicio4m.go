package main

import (
	"errors"
	"fmt"
)

// Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
// requiriendo calcular el mínimo, máximo y promedio.
//
// Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
// máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
// definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
// indicamos en la función anterior

func studentCalc() (func(numbers ...int) float64, error) {
	var op string
	fmt.Println("Indique que tipo de calculo quiere realizar(min, max o avg)")
	fmt.Scan(&op)

	switch op {
	case "min":
		return min, nil
	case "max":
		return max, nil
	case "avg":
		return avg, nil
	default:
		return nil, errors.New("No existe ese calculo")
	}
}

func min(numbers ...int) float64 {
	var min int = numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return float64(min)
}

func max(numbers ...int) float64 {
	var max int = numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return float64(max)
}

func avg(numbers ...int) float64 {
	var accum float64
	for _, number := range numbers {
		accum += float64(number)
	}
	return accum / float64(len(numbers))
}

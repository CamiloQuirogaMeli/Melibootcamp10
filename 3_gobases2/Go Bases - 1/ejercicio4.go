package main

import (
	"errors"
	"fmt"
)

/*
	Ejercicio 4 - Calcular estadísticas
	Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
	requiriendo calcular el mínimo, máximo y promedio.

	Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
	máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
	definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
	indicamos en la función anterior
*/

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func ejecutarAccion(accion string) (func(...float64) (float64, error), error) {
	switch accion {
	case minimo:
		return Ominimo, nil
	case promedio:
		return Opromedio, nil
	case maximo:
		return Omaximo, nil
	}
	return nil, errors.New("Accion no existe")
}

func Ominimo(valores ...float64) (float64, error) {
	min := valores[0]
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		if valor < min {
			min = valor
		}
	}
	return min, nil
}

func Opromedio(valores ...float64) (float64, error) {
	var suma, promedio = 0.0, 0.0
	contador := 0.0
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		suma += valor
		contador++
	}
	promedio = suma / contador
	return promedio, nil
}

func Omaximo(valores ...float64) (float64, error) {
	max := 0.0
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		if valor > max {
			max = valor
		}
	}
	return max, nil
}

func main() {
	accion := ""
	fmt.Print("Digite la funcion que desea:\nminimo\nmaximo\npromedio\n")
	fmt.Scan(&accion)

	mifuncion, err := ejecutarAccion(accion)
	if err != nil {
		fmt.Println(err)
	} else {
		respuesta, errors := mifuncion(-2, 3, 4, 5, 6, 7, 8, 10)
		if errors != nil {
			fmt.Println(errors)
		} else {
			fmt.Println(respuesta)
		}
	}
}

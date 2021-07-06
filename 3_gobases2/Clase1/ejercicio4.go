package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	/*
		Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
		requiriendo calcular el mínimo, máximo y promedio.
		Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
		máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
		definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
		indicamos en la función anterior

	*/
	var operacion string
	notas := make([]float64, 0)
	var entrada string
	fmt.Println("Digita el tipo de operación a realizar:")
	fmt.Scanln(&operacion)

	fmt.Println("Escribe las notas del alumno")
	fmt.Println("(siguiente nota presiona enter)")
	fmt.Println("(Para el promedio escribe p)")
	fmt.Scanln(&entrada)
	for entrada != "p" {

		nota, err := strconv.ParseFloat(entrada, 8)
		if err != nil {
			fmt.Println("Error, dato no válido")
			fmt.Scanln(&entrada)

		} else {
			notas = append(notas, nota)
			fmt.Scanln(&entrada)
		}
	}

	funcion, err := calcularEstadisticas(operacion)
	if err != nil {
		fmt.Println("operación no válida")
		fmt.Println("Las operaciones disponibles son:")
		fmt.Println("minimo \nmaximo \npromedio")
	} else {
		resultado, err := funcion(notas)
		if err != nil {
			fmt.Println("Datos erroneos o vacios")
		} else {
			fmt.Print("El resultado es:")
			fmt.Println(resultado)

		}
	}

}

func calcularEstadisticas(operacion string) (func([]float64) (float64, error), error) {
	switch operacion {
	case minimo:
		return calcularMinimo, nil
	case maximo:
		return calcularMaximo, nil
	case promedio:
		return calcularPromedio, nil
	}
	return nil, errors.New("operación no válida")
}

func calcularMinimo(notas []float64) (float64, error) {
	var resultado float64
	if notas == nil {
		return 0, errors.New("No se digitaron los datos")
	}

	for i, nota := range notas {
		if i == 0 {
			resultado = nota
		} else {
			if nota < resultado {
				resultado = nota
			}
		}
	}
	return resultado, nil
}

func calcularMaximo(notas []float64) (float64, error) {
	var resultado float64
	if notas == nil {
		return 0, errors.New("No se digitaron los datos")
	}

	for i, nota := range notas {
		if i == 0 {
			resultado = nota
		} else {
			if nota > resultado {
				resultado = nota
			}
		}
	}
	return resultado, nil
}

func calcularPromedio(notas []float64) (float64, error) {
	var resultado float64
	if notas == nil {
		return 0, errors.New("No se digitaron los datos")
	}

	for _, nota := range notas {
		resultado += nota
	}
	numeroNotas := float64(len(notas))
	resultado = resultado / numeroNotas

	return resultado, nil
}

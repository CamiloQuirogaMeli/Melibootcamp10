package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/*
		Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
		Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
		devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
	*/
	notas := make([]float64, 0)
	var entrada string
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
	fmt.Println("El promedio es el siguiente:")
	promedio, err := CalcularPromedio(notas)
	if err != nil {
		fmt.Println("Error, promedio no calculado")
	}
	fmt.Println(promedio)
}

func CalcularPromedio(notas []float64) (float64, error) {
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

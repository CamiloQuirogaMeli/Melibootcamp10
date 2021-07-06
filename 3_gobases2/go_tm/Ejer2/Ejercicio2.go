/*
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	resultado, errorResultado := generarPromedio(5.5, 2.2, 3.3, 4.4, -5)

	if errorResultado != nil {
		fmt.Println(errorResultado)
		os.Exit(0)
	}

	println("El resultado de las notas promediadas del estudiante son: ", resultado)
}

func generarPromedio(notas ...float64) (float64, error) {
	var resultado float64
	println("Las notas ingresadas, son las siguientes: ")
	for _, value := range notas {
		println(strconv.FormatFloat(value, 'f', -2, 64))
		if math.Signbit(value) {
			return 0, errors.New("Existe una nota negativa")
		}
		resultado += value
	}

	return resultado / float64(len(notas)), nil

}

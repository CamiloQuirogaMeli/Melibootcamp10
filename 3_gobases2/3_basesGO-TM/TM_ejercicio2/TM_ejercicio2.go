package main

import (
	"errors"
	"fmt"
)

/*

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

*/

/*Calculo el promedio de calificaciones*/
/*func calculaPromedio(calificaciones ...int) (float64, error) {*/
func calculaPromedio(calificaciones []int) (float64, error) {

	/*Cantidad de notas del alumno*/
	N := len(calificaciones)

	/*Calculo el promedio*/
	promedio := 0.0

	for _, nota := range calificaciones {

		if float64(nota) < 0 {
			return 0.0, errors.New("la calificación no puede ser negativa")
		}

		promedio += float64(nota)

	}

	promedio = promedio / float64(N)

	return promedio, nil

}

func main() {

	var calificaciones = []int{10, 10, 10, 9, 8, 9, 8, 7}

	promedio, err := calculaPromedio(calificaciones)

	fmt.Println("Calificaciones: ", calificaciones)
	fmt.Println("Promedio: ", promedio)
	fmt.Println("Errores en la operación: ", err)

}

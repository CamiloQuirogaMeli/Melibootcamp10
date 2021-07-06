package main

import (
	"errors"
	"fmt"
)

//
// Ejercicio 2 - Calcular promedio
// Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
// devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
//

func promedio(notas ...int) (float64, error) {
	var promedio int
	var result float64
	for _, nota := range notas {
		if nota > 0 {
			promedio += nota
		} else {
			return 0, errors.New("Ocurrió un error, no puede contener notas menores a 0")
		}
	}
	result = float64(promedio) / float64(len(notas))
	return result, nil
}

func main() {

	fmt.Println("===== Calcular promedio =====")

	fmt.Println("== Estudiante 1 ==")

	fmt.Println("Notas: 10, 50, -1, 20")
	promdEst1, errEst1 := promedio(10, 50, -1, 20)

	if errEst1 != nil {
		fmt.Println(errEst1)
	} else {
		fmt.Println("Promedio calculado:", promdEst1)
	}

	fmt.Println("== Estudiante 2 ==")
	fmt.Println("Notas: 100, 15, 50, 25, 60, 35")
	promdEst2, errEst2 := promedio(100, 15, 50, 25, 60, 35)

	if errEst2 != nil {
		fmt.Println(errEst2)
	} else {
		fmt.Println("Promedio calculado:", promdEst2)
	}

}

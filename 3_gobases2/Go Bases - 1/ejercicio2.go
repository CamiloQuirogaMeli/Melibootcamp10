package main

import (
	"errors"
	"fmt"
)

/*
	Ejercicio 2 - Calcular promedio
	Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos devuelva
	el promedio y un error en caso que uno de los números ingresados sea negativo
*/

func promedioNotas(notas ...float64) (float64, error) {
	var suma, resultado = 0.0, 0.0
	contador := 0
	for i, nota := range notas {
		if notas[i] < 0 {
			return 0, errors.New("Existe una nota negativa")
		}
		suma += nota
		contador++
	}
	resultado = suma / float64(contador)
	return resultado, nil
}

func main() {
	var notas, err = promedioNotas(0.1, 4.0, 5.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", notas)
	}
}

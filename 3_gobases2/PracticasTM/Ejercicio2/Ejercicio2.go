package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Ejercicio 2 - Calcular promedio
		Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
		Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
		devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
	*/

	prom, err := promedio(50, 60, 10, 80, -45)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es", prom)
	}
}

func promedio(notas ...int) (int, error) {
	var sum int
	cant := len(notas)
	for _, value := range notas {
		sum += value
		if value < 0 {
			return 0, errors.New("Se ha ingresado un numero negativo")
		}
	}

	return sum / cant, nil
}

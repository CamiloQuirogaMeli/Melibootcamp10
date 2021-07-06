package main

import (
	"errors"
	"fmt"
)

/*
Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
requiriendo calcular el mínimo, máximo y promedio

Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior
Ejemplo:

const (
minimo = "minimo"
promedio = "promedio"
maximo = "maximo"
)

...
minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...
valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

*/

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

/*Chequeo que las calificaciones sean válidas*/
func checkNotas(calificaciones []int) bool {
	for _, nota := range calificaciones {
		if nota < 0 {
			return false
		}
	}
	return true
}

/*Menu de acciones posibles*/
func operacion(op string) (func(calificaciones []int) float64, error) {

	switch op {
	case MINIMO:
		return calculaMinimo, nil
	case PROMEDIO:
		return calculaPromedio, nil
	case MAXIMO:
		return calculaMaximo, nil
	default:
		return calculaMinimo, errors.New("la operación no es válida")

	}

}

/*Calculo la calificación mínima*/
func calculaMinimo(calificaciones []int) float64 {

	/*Asigno el primer elemento y comienzo el calculo*/
	min := calificaciones[0]

	for _, nota := range calificaciones {

		if nota < min {
			min = nota
		}

	}

	return float64(min)
}

/*Calculo la calificacion máxima*/
func calculaMaximo(calificaciones []int) float64 {

	/*Asigno el primer elemento y comienzo el calculo*/
	max := calificaciones[0]

	for _, nota := range calificaciones {

		if nota > max {
			max = nota
		}

	}

	return float64(max)
}

/*Calculo el promedio de calificaciones*/
func calculaPromedio(calificaciones []int) float64 {

	/*Cantidad de notas del alumno*/
	N := len(calificaciones)

	/*Calculo el promedio*/
	promedio := 0.0

	for _, nota := range calificaciones {

		promedio += float64(nota)

	}

	promedio = promedio / float64(N)

	return promedio

}

func main() {

	var calificaciones = []int{10, 10, 9, 9, 8, 8, 7}

	/*Chequeo calificaciones*/

	if checkNotas(calificaciones) {
		var valor float64

		/*Calculo la máxima calificación del alumno y la muestro*/
		funcion, err := operacion(MAXIMO)
		/* En este caso dado que yo ingreso la operación, no tiene demasiado sentido,
		pero si se ingresase por consola, verifico que sea válida*/
		if err == nil {
			valor = funcion(calificaciones)
		}

		fmt.Println("Calificaciones: ", calificaciones)
		fmt.Println("Máxima: ", valor)

		/*Idem anterior para mínimo*/
		funcion, err = operacion(MINIMO)
		if err == nil {
			valor = funcion(calificaciones)
		}

		fmt.Println("Mínima: ", valor)

		/*Idem para promedio*/
		funcion, err = operacion(PROMEDIO)
		if err == nil {
			valor = funcion(calificaciones)
		}

		fmt.Println("Promedio: ", valor)
	} else {
		fmt.Println("Las calificaciones no pueden ser negativas")
	}

}

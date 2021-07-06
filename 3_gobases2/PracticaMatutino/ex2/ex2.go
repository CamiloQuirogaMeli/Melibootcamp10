package main

import (
	"errors"
	"fmt"
)

func main() {

	var avg, err = getAvg(5.4, 2.0, 5.5, 4.4, 3.0, 8.1, 9.2, 1.1)

	if err == nil {
		fmt.Println(avg)
	} else {
		fmt.Println(err)
	}

}

func getAvg(califications ...float64) (average float64, err error) {

	sum := 0.0

	if len(califications) == 0 {
		return
	}

	for i := range califications {
		if califications[i] < 1 {
			err = errors.New("se encontro una calificacion negativa")
			break
		}
		sum += califications[i]
	}
	average = sum / (float64)(len(califications))

	return
}

/*
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

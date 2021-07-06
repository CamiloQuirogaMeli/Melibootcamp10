package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func ejercicio4() {

	minFunc, err := operacion(minimo)
	if err != nil {
		fmt.Println(err)
	}

	promFunc, err := operacion(promedio)
	if err != nil {
		fmt.Println(err)
	}
	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Println(err)
	}

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("El valor minimo es: %v\nEl valor promedio es: %v\nEl Valor maximo es: %v\n", valorMinimo, valorPromedio, valorMaximo)

}

func operacion(operacion string) (func(...float64) float64, error) {
	switch operacion {
	case minimo:
		return minFunc, nil
	case promedio:
		return funcProm, nil
	case maximo:
		return maxFunc, nil
	default:
		return nil, errors.New("funcion no definida")
	}
}

func minFunc(calificaciones ...float64) (min float64) {
	min = calificaciones[0]
	for _, calificacion := range calificaciones {

		if calificacion < min {
			min = calificacion
		}
	}

	return
}

func funcProm(calificaciones ...float64) (promedio float64) {

	var suma float64 = 0

	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0
		}
		suma += calificacion
	}

	promedio = (suma / float64(len(calificaciones)))

	return

}

func maxFunc(calificaciones ...float64) (max float64) {
	max = calificaciones[0]
	for _, calificacion := range calificaciones {

		if calificacion > max {
			max = calificacion
		}
	}

	return
}

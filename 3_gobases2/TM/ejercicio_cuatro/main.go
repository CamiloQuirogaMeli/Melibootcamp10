package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	MAXIMO   = "maximo"
	PROMEDIO = "promedio"
)

func calMinimo(notas ...float32) (result float32) {

	result = 0
	for i, nota := range notas {
		if i == 0 {
			result = nota
		}
		if nota < result {
			result = nota
		}
	}
	return
}

func calMaximo(notas ...float32) (result float32) {

	result = 0
	for _, nota := range notas {
		if nota > result {
			result = nota
		}
	}
	return
}

func calPromedio(notas ...float32) (result float32) {

	result = 0
	var suma float32
	for _, nota := range notas {
		suma += nota
	}
	result = suma / float32(len(notas))
	return
}

func operacion(operacionType string) (funcion func(...float32) float32, err error) {
	err = nil
	switch operacionType {
	case MAXIMO:
		funcion = calMaximo
	case MINIMO:
		funcion = calMinimo
	case PROMEDIO:
		funcion = calPromedio
	default:
		err = errors.New("opcion incorrecta")
	}
	return
}

func main() {

	minFunc, minErr := operacion(MINIMO)
	propFunc, propErr := operacion(PROMEDIO)
	maxFunc, maxErr := operacion(MAXIMO)

	if minErr == nil {
		valorMinimo := minFunc(2, 3, 4)
		fmt.Println("valor minimo:", valorMinimo)
	}

	if propErr == nil {
		valorProm := propFunc(2, 3, 4)
		fmt.Println("valor promedio:", valorProm)
	}

	if maxErr == nil {
		valorMax := maxFunc(2, 3, 4)
		fmt.Println("valor maximo:", valorMax)
	}

}

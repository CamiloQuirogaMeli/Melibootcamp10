package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func operacion(funcEstad string) (func(valores ...float64) float64, error) {
	switch funcEstad {
	case MINIMO:
		return minFunc, nil
	case PROMEDIO:
		return promFunc, nil
	case MAXIMO:
		return maxFunc, nil
	}
	msg := errors.New("La operacion no está definida")
	return nil, msg
}

func minFunc(valores ...float64) float64 {
	numeroMinimo := valores[0]
	for _, value := range valores {
		if value < numeroMinimo {
			numeroMinimo = value
		}
	}
	fmt.Println("el minimo es: ", numeroMinimo)
	return numeroMinimo
}
func promFunc(valores ...float64) float64 {
	var prom float64
	var contador float64 = 0
	var promedio float64
	for _, value := range valores {
		prom += value
		contador++
	}
	promedio = prom / contador
	fmt.Println("el promedio es: ", promedio)
	return promedio
}

func maxFunc(valores ...float64) float64 {
	numeroMaximo := valores[0]
	for _, value := range valores {
		if value > numeroMaximo {
			numeroMaximo = value
		}
	}
	fmt.Println("el maximo es: ", numeroMaximo)
	return numeroMaximo
}

package main

import "errors"

func promedio(values ...float64) (float64, error) {

	var prom float64
	var contador float64 = 0
	var promedio float64

	for _, value := range values {
		if value < 0 {
			return 0, errors.New("Hay una nota negativa")

		} else if value > 0 {
			prom += value
			contador++
		}
		promedio = prom / contador
	}

	return promedio, nil

}

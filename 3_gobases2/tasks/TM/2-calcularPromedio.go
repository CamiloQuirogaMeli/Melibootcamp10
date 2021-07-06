package main

import "errors"

func calcularPromedio(notas []float64) (media float64, err error) {
	var sumatoria float64 = 0

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Las notas no pueden ser negativas")
		}
		sumatoria += value
	}

	media = sumatoria / float64(len(notas))

	return media, nil
}

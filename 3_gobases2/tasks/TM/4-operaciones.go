package main

import "errors"

func opMin(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("No hay registros para medir")
	}

	var min float64 = 0
	// flo := 0.0

	for i, value := range notas {
		if value < 0 {
			return 0, errors.New("La nota no puede ser negativa")
		}

		if i == 0 || value < min {
			min = value
		}
	}

	return min, nil
}

func opMax(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("No hay registros para medir")
	}

	var max float64 = 0

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("La nota no puede ser negativa")
		}

		if max < value {
			max = value
		}
	}

	return max, nil
}

func opPromedio(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("No hay registros para medir")
	}

	var sumatoria float64 = 0

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("La nota no puede ser negativa")
		}

		sumatoria = sumatoria + value
	}

	var media float64 = sumatoria / float64(len(notas))

	return media, nil
}

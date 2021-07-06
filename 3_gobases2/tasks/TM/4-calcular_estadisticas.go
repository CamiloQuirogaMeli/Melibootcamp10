package main

import (
	"errors"
)

func calcularEstadisticas(tipoDeCalculo string) (func(notas ...float64) (float64, error), error) {
	switch tipoDeCalculo {
	case "minimo":
		return opMin, nil
	case "maximo":
		return opMax, nil
	case "promedio":
		return opPromedio, nil
	}

	errManagement := func(notas ...float64) (float64, error) { return 0, nil }

	return errManagement, errors.New("error: tipo de cálculo no encontrado")
}

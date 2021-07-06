package main

import "errors"

func calcularSalario(minutos int, categoria string) (salary int, err error) {
	switch categoria {
	case "A":
		salary = minutos * 3000
		return salary, nil
	case "B":
		salary = minutos * 1500
		return salary, nil
	case "C":
		salary = minutos * 1500
		return salary, nil
	default:
		return 0, errors.New("error: categoría no encontrada")
	}
}

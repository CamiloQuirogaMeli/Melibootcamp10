package main

import "errors"

func calcularPromedio(notas ...int) (promedio float64, err error) {
	var sumaTotal int
	for _, valor := range notas {
		if valor < 0 {
			err = errors.New("ningún valor puede ser menor a 0")
		} else {
			sumaTotal += valor
		}
	}

	promedio = float64(sumaTotal) / float64(len(notas))

	return
}

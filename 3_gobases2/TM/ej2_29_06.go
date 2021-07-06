package main

import (
	"errors"
	"fmt"
)

func promedio(valores []float64) (float64, error) {
	var resultado float64

	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("la calificacion no puede ser negativa")
		} else {
			resultado += valor
		}
	}
	return resultado / float64(len(valores)), nil
}

func main() {

	s := make([]float64, 0)
	for {
		fmt.Println("Ingrese calificaciones")
		var i float64
		fmt.Scan(&i)
		if i < 0 {
			break
		}
		s = append(s, i)
		a, err := promedio(s)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Las calificaciones son:", s)
		fmt.Println("El promedio es:", a)
	}
}

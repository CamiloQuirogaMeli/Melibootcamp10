package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...int) (float64, error) {
	sum := 0

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("no se puede ingresar una nota negativa")
		} else {
			sum += nota
		}
	}
	return float64(sum) / float64(len(notas)), nil
}

func main() {
	res, err := calcularPromedio(-2, 3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de las notas es: %.2f\n", res)
	}
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := calcularPromedio(1, 2, 3, -4, 5)

	if err != nil {
		fmt.Println("Hubo un error: ", err)
	}

	fmt.Println("El promedio es: ", prom)
}

func calcularPromedio(valores ...float64) (float64, error) {
	var sum float64
	for _, value := range valores {
		if value < 0 {
			return 0, errors.New("Existe un valor negativo")
		}

		sum += value
	}

	return sum / float64(len(valores)), nil
}

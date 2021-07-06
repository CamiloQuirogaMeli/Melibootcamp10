package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedio(-2, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es ", promedio)
	}
}

func calcularPromedio(valores ...float32) (float32, error) {
	var sumatoria float32 = 0.0
	for _, value := range valores {
		if value < 0 {
			return 0, errors.New("Número negativo")
		}
		sumatoria += value
	}

	return sumatoria / float32(len(valores)), nil
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := promedio(5, 2, 3, 4, 10, 8)
	//prom, err := promedio(5, -2, 3, 4, 10, 8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", prom)
	}
}

func promedio(values ...int) (float64, error) {
	var sumaTotal, cantNotas int = 0, 0
	var resultado float64

	for _, value := range values {
		if value < 0 {
			return 0, errors.New("No puede haber una nota menor a 0")
		} else {
			sumaTotal += value
			cantNotas++
		}
	}
	resultado = float64(sumaTotal) / float64(cantNotas)
	return resultado, nil
}

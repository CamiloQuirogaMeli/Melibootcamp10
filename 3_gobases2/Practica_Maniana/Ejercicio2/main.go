package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (float64, error) {
	var cant, suma int
	var promedio float64
	for _, values := range notas {
		if values < 0 {
			return 0, errors.New("Error numero negativo")
		} else {
			cant++
			suma = suma + values
		}
	}
	promedio = float64(suma) / float64(cant)

	return promedio, nil
}

func main() {
	p, err := promedio(5, 5)

	if err != nil {
		fmt.Println("Hubo un problema: ", err)
	} else {
		fmt.Println("Promedio: ", p)
	}
}

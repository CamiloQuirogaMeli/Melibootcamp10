package main

import (
	"errors"
	"fmt"
)

func main() {

	promedio, err := calcularPromedio(2, 6, 7, 4, 3, 7.25, 8, 5, 9.5, 8, 8, 4)

	if err != nil {
		fmt.Printf("%s \n", err)
	} else {
		fmt.Printf("El promedio es %f \n", promedio)
	}

}

func calcularPromedio(notas ...float32) (float32, error) {
	var count float32 = 0.0

	for _, val := range notas {
		if val < 0 {
			return -1, errors.New("Se encontro una nota negativa ")
		}

		count += val
	}

	return count / float32(len(notas)), nil

}

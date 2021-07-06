package main

import (
	"errors"
	"fmt"
)

func main() {

	notas := []int{10, 3, 5, 7, 8, 10}
	avd, err := avdCalc(notas)
	if err != nil {
		fmt.Print(err)
	} else {

		fmt.Printf("El promedio es %.2f\n", avd)
	}
}

func avdCalc(notas []int) (float32, error) {
	sumatoria := 0.0
	for _, nota := range notas {
		if nota < 0 {
			return 0.0, errors.New("No se puede ingresar una nota negativa\n")
		}
		sumatoria += float64(nota)
	}
	return float32(sumatoria / float64(len(notas))), nil

}

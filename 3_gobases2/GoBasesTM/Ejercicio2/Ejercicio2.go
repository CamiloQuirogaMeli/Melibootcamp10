package main

import (
	"errors"
	"fmt"
)

func calculoPromedio(valores ...int) (float32, error) {

	var suma int = 0
	var c int = 0

	for _, value := range valores {
		if value < 0 {
			return 0.0, errors.New("La nota no puede ser negativa")
		} else {
			c++
			suma += value
		}
	}

	return float32(suma / c), nil
}

func main() {
	var sliceUsr = []int{8, 6, 10, 4, 3, 2}

	prom, e := calculoPromedio(sliceUsr...)

	if e == nil {
		fmt.Printf("Promedio: %.2f\n", prom)
	} else {
		fmt.Printf("Ha ocurrido el siguiente error: %s\n", e.Error())
	}

}

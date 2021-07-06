package main

import (
	"errors"
	"fmt"
)

func main() {
	avg, err := average(4, 6, 4, 6, 4, 6, -1)

	if err == nil {
		fmt.Println("El promedio de las notas es:", avg)
	} else {
		fmt.Println(err)
	}

}

func average(calific ...float64) (float64, error) {
	var count int = 0
	var suma float64 = 0

	for _, value := range calific {
		if value > 0 {
			suma += value
			count++
		} else if value < 0 {
			return 0, errors.New("ERROR! Se cargó un numero negativo")
		}
	}

	return suma / float64(count), nil
}

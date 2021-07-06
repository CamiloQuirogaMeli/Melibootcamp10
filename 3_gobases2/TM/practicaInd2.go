package main

import (
	"errors"
	"fmt"
)

func main() {

	prom, err := getAvg(8, 9, 5, 3, 2, 6, -4)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio general es %f \n", prom)
	}

}

func getAvg(note ...float64) (float64, error) {

	var prom, sum float64
	i := 0

	for _, v := range note {

		if v <= 0 {
			return 0, errors.New("La nota mínima es 1. Vuelva a ingresar los valores")
		}
		sum += v
		i++

	}

	prom = sum / float64(i)

	return prom, nil
}

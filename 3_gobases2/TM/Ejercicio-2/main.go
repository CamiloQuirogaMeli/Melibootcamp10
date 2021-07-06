package main

import (
	"errors"
	"fmt"
)

func calcAvg(values ...float64) (float64, error) {
	acumGrades := 0.0
	for _, value := range values {
		if value < 0.0 {
			return 0.0, errors.New("una calificacion no puede ser menor a 0")
		}
		acumGrades += value
	}
	acumGrades = acumGrades / float64(len(values))
	return acumGrades, nil
}

func main() {
	var s = []float64{9.5, 10, 4, 5, 6, 7, 8, -5}
	fmt.Println("Dadas las notas:", s)
	res, err := calcAvg(9.5, 10, 4, 5, 6, 7, 8, -5)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Printf("Su promedio es %.2f", res)
	}

}

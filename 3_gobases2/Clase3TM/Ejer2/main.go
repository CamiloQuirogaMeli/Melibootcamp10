package main

import (
	"errors"
	"fmt"
)

func main() {
	avgScore, err := averageScore(1, -5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("La nota promedio es de  %.1f\n", avgScore)
	}
}

func averageScore(grades ...float64) (float64, error) {
	avgGrade := 0.0
	for _, grade := range grades {
		if grade < 0 {
			return avgGrade, errors.New("Todas las notas deben ser mayores o iguales a 0")
		} else {
			avgGrade += grade
		}
	}
	avgGrade = avgGrade / float64(len(grades))
	return avgGrade, nil
}

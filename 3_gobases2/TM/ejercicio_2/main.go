package main

import (
	"fmt"
	"errors"
)

func containsNegative(grades []float32) bool {
	for _, grade := range grades {
		if grade < 0 {
			return true
		}
	}

	return false
}

func gradeAverage(grades ...float32) (float32, error) {
	if containsNegative(grades) {
		return -1, errors.New("Una calificacion no puede ser negativa")
	}

	var total float32 = 0
	for _, grade := range grades {
		total += grade
	}
	
	gradesCount := len(grades)
	average := total / float32(gradesCount)

	return average, nil
}

func main() {

	fmt.Println("* Caso 1: calificaciones positivas")
	validGrades := []float32{10, 6, 8, 1, 6}
	fmt.Println("\tCalificaciones: ", validGrades)
	average, _ := gradeAverage(validGrades...)
	fmt.Println("\tPromedio: ", average)

	fmt.Println("* Caso 2: calificaciones negativas")
	invalidGrades := []float32{10, 6, 8, -1, 6}
	fmt.Println("\tCalificaciones: ", invalidGrades)
	_, err := gradeAverage(invalidGrades...)
	fmt.Println("\tError: ", err)
}

package main

import (
	"errors"
	"fmt"
)

func average(grades ...int) (float64, error) {
	gradeSum, students := 0, 0
	for _, grade := range grades {
		if grade < 0 {
			return 0.0, errors.New("No puede haber notas negativas.")
		}
		gradeSum += grade
		students++
	}
	return float64(gradeSum) / float64(students), nil
}

func main() {
	var avg, _ = average(1, 4, 8, 3, 2, 6, 7, 4, 9, 10, 5)
	fmt.Printf("El promedio de notas es %.2f\n", avg)
}

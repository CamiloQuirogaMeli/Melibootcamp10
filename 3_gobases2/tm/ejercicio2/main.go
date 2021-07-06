package main

import (
	"errors"
	"fmt"
	"os"
)

func avgGrades(grades ...float64) (avg float64, err error) {
	var sumGrades float64
	for _, grade := range grades {
		if grade < 0 {
			err = errors.New("the student has invalid grades")
			return 0, err
		}
		sumGrades += grade
	}
	avg = sumGrades / float64(len(grades))
	return avg, nil
}

func main() {
	// Por alumno...
	var grades = []float64{4.5, -10, 9, 7, 6, 5.5, 2, 1}
	avg, err := avgGrades(grades...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(avg)
}

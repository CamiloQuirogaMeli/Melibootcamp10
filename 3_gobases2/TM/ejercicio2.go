package main

import (
	"errors"
	"fmt"
)

func main() {
	average, err := getAverage(5, 5, 5, -5, 10, 10, 10, 10)

	if err == nil {
		fmt.Printf("El promedio es: %.2f\n", average)
	} else {
		fmt.Println(err)
	}
}

func getAverage(ratings ...int) (float64, error) {
	var ratingsSum float64
	for _, rating := range ratings {
		if rating >= 0 {
			ratingsSum += float64(rating)
		} else {
			return 0, errors.New("Calificacion Negativa")
		}
	}

	return ratingsSum / float64(len(ratings)), nil
}

package main

import (
	"errors"
	"fmt"
)

func average(grades ...int) (float64, error) {
	sum := 0
	for _, v := range grades {
		if v < 0 {
			return 0, errors.New("Las calificaciones no pueden ser negativas")
		}
		sum += v
	}
	avg := float64(sum) / float64(len(grades))

	return avg, nil
}

func main() {
	grades := []int{1, 2, 3, 4, 5}

	if avg, err := average(grades...); err == nil {
		fmt.Println("El promedio de las calificaciones es", avg)
	} else {
		fmt.Println(err)
	}
}

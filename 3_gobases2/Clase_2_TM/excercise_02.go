package main

import (
	"errors"
	"fmt"
)

func main() {
	avg, err := avg(7.00, 7.00, 7.50, 7.50, 4.00)

	fmt.Println(avg, err)
}

func avg(califications ...float64) (float64, error) {
	var avg, sum float64
	for _, calification := range califications {
		if calification < 0 {
			return 0, errors.New("A calification is negative.")
		}
		sum += calification
	}
	avg = divide(sum, float64(len(califications)))
	return avg, nil
}

func divide(sumCalifications float64, countCalifications float64) float64 {
	return sumCalifications / countCalifications
}

package main

import (
	"errors"
	"fmt"
)

func main() {

	mean, err := getMean(5, 5, 5, 5, 5, 0)
	if err == nil {
		fmt.Printf("Mean of grades is equal to %.2f", mean)
	} else {
		fmt.Println(err)
	}

}

func getMean(nums ...float64) (float64, error) {

	mean := 0.0

	for _, grade := range nums {
		if grade < 0 {
			return 0, errors.New("grades can't be negative")
		}
		mean += grade
	}
	return mean / float64(len(nums)), nil
}

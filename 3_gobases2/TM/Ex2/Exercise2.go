package main

import (
	"errors"
	"fmt"
)

func main() {
	notes := []int{5, 5, 5}

	res, err := calcAvg(notes...)

	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("Average Calification:", res)
	}

}

func calcAvg(califications ...int) (float64, error) {

	if len(califications) == 0 {
		return 0, nil
	}

	var sum int

	for _, note := range califications {
		if note < 0 {
			return 0, errors.New("A calification cant be negative")
		} else {
			sum += note
		}
	}

	return (float64(sum) / float64(len(califications))), nil

}

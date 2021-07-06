package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := averageNotes(10, 8, -1, 4)
	fmt.Println(result, err)
}

func averageNotes(notes ...int) (float64, error) {
	var total float64 = 0
	for _, value := range notes {
		if value < 0 {
			return float64(value), errors.New("Hay un valor negativo")
		}
		total += float64(value)
	}
	return total / float64(len(notes)), nil
}

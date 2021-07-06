package main

import (
	f "fmt"
	e "errors"
)

func average(values ...float32) (float32, error) {
	var sum float32
	for _, value := range values {
		if value < 0 {
			return 0, e.New("Negative number")
		} else {
			sum += value
		}
	}
	return sum / float32(len(values)), nil
}

func main() {
	res, err := average(-8, 8, 8, 8, 8, 10, 10, 10, 10)
	if err == nil {
		f.Printf("%.2f\n", res)
	} else {
		f.Println(err)
	}
}
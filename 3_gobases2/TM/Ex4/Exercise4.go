package main

import (
	"errors"
	"fmt"
)

const (
	MIN = "min"
	MAX = "max"
	AVG = "avg"
)

func main() {

	numbers := []int{1, 2, 3, 4}

	opAvg, _ := calculate(AVG)
	opMin, _ := calculate(MIN)
	opMax, _ := calculate(MAX)
	_, err := calculate("asd")

	fmt.Println("Average:", opAvg(numbers...))
	fmt.Println("Min:", opMin(numbers...))
	fmt.Println("Max:", opMax(numbers...))
	fmt.Println("Error:", err)
}

func calculate(operation string) (func(numbers ...int) float64, error) {

	switch operation {
	case AVG:
		return calcAvg, nil
	case MAX:
		return calcMax, nil
	case MIN:
		return calcMin, nil
	default:
		return nil, errors.New("invalid option")
	}

}

func calcAvg(numbers ...int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	var sum int

	for _, number := range numbers {
		sum += number
	}

	return (float64(sum) / float64(len(numbers)))

}

func calcMax(numbers ...int) float64 {

	var max int = numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return float64(max)

}

func calcMin(numbers ...int) float64 {

	var min int = numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return float64(min)

}

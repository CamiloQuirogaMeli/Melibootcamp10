package main

import (
	er "errors"
	f "fmt"
)

const (
	minimum = "min"
	mean    = "mean"
	maximum = "max"
)

func main() {

	minFunc, err1 := calculateStatistics(minimum)
	meanFunc, err2 := calculateStatistics(mean)
	maxFunc, err3 := calculateStatistics(maximum)

	handleErrors(err1)
	handleErrors(err2)
	handleErrors(err3)

	min_value := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	f.Println("Minimum: ", min_value)

	max_value := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	f.Println("Maximum: ", max_value)

	mean_value := meanFunc(2, 3, 3, 4, 1, 2, 4, 5)
	f.Println("Mean: ", mean_value)
}

func calculateMin(notes ...float32) float32 {

	var min float32 = notes[0]

	for _, value := range notes {
		if value < min {
			min = value
		}
	}

	return min
}

func calculateMax(notes ...float32) float32 {
	var max float32 = notes[0]

	for _, value := range notes {
		if value > max {
			max = value
		}
	}

	return max
}

func calculateMean(notes ...float32) float32 {

	var sum float32 = 0

	for _, value := range notes {
		sum += value
	}

	return sum / float32(len(notes))
}

func calculateStatistics(calculation string) (func(notes ...float32) float32, error) {

	switch calculation {

	case minimum:
		return calculateMin, nil

	case mean:
		return calculateMean, nil

	case maximum:
		return calculateMax, nil

	default:
		return nil, er.New("This calculation is not defined")
	}

}

func handleErrors(err error) {
	if err != nil {
		f.Println(err)
	}
}

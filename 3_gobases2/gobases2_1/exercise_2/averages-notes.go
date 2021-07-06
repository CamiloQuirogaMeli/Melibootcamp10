package main

import f "fmt"

func main() {
	f.Println(calculateAverage(1, 4, 5))
}

func calculateAverage(notes ...float32) float32 {

	var sum float32
	var result float32

	for _, value := range notes {
		sum += value
	}

	result = sum / float32(len(notes))

	return result
}

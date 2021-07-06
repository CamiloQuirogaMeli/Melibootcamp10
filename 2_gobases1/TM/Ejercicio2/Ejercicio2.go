package main

import "fmt"

func main() {
	var (
		temperature float32 = 32.1
		humidity    float32 = 99.5
		airPressure float32 = 1013.25
	)

	fmt.Println("Temperature: ", temperature)
	fmt.Println("Humidity: ", humidity)
	fmt.Println("Air Pressure: ", airPressure)
}

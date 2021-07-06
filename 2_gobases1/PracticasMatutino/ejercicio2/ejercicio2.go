package main

import "fmt"

var (
	temperature int
	humidity    int
	pressure    float64
)

func main() {
	temperature = 4
	humidity = 10
	pressure = 10.2

	fmt.Println("Temperature: ", temperature)
	fmt.Println("Humidity: ", humidity)
	fmt.Println("Pressure: ", pressure)
}

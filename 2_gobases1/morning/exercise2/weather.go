package main

import "fmt"

func main() {

	//I would asign float types for all variables, because they can range in a continous interval.
	var temperature float64 = 24.5
	var humidity float32 = 70
	var pressure float64 = 700

	fmt.Println("Humidity percentage:", humidity, "%\nCurrent temperature:", temperature, "degrees\nPressure:", pressure, "mm")
}

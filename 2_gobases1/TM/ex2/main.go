package main

import (
	"fmt"
)

func main() {
	var (
		temperature float64 = 19
		humidity    float64 = 0.52
		pressure    float64 = 753
	)
	fmt.Println("Welcome!")
	fmt.Println("Temperature:", temperature, "C")
	fmt.Println("Humidity:", humidity*100, "%")
	fmt.Println("Pressure:", pressure, "hPa")
}

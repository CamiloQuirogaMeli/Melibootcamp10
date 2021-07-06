package main

import "fmt"

func main() {
	var (
		temperature       int     = 18
		relative_humidity float32 = 0.57
		precipitation     string  = "0.1"
	)

	fmt.Printf("Temperature in Bogotá: %d ºC \n", temperature)
	fmt.Printf("Relative Humidity in Bogotá: %.2f mm \n", relative_humidity)
	fmt.Printf("Precipitation: %s \n", precipitation)
}

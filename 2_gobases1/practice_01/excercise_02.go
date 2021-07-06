package main

import "fmt"

const (
	DEGREES_CELSISUS = "ºC"
	PERCENTAGE       = "%"
	HPA              = "hPa"
)

func main() {

	var (
		temperature = 26.9
		humidity    = 50
		pressure    = 1008.4
	)

	fmt.Println(fmt.Sprint(temperature, DEGREES_CELSISUS), fmt.Sprint(humidity, PERCENTAGE), fmt.Sprint(pressure, HPA))

}

package main

import "fmt"

var (
	temperature          int8   = 29
	humidity             uint8  = 21
	atmospheric_pressure uint16 = 1015
)

func main() {
	fmt.Printf(
		`MEASUREMENT VALUES NEIVA
		Temperature = %v ºC
		Humidity = %v %%
		Atmospheric_pressure = %v hPa`, temperature, humidity, atmospheric_pressure)
}

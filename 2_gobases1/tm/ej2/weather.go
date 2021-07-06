package ej2

import "fmt"

func Weather()  {
	var (
		temperature = 21.0
		pressure = 2640.0
		humidity = 0.55
	) //float
	fmt.Printf("Temperature = %.1f\nPressure = %.1f\nHumidity = %.2f\n",temperature,pressure,humidity)
}
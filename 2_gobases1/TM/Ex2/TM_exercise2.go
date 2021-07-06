package main

import (
	"fmt"
)

func main() {
	var tempCelsius int = 20
	var humPercent int = 36
	var presHPa float32 = 1009.5

	fmt.Println("Temperature:", tempCelsius, "ºC")
	fmt.Println("Humidity:", humPercent, "%")
	fmt.Println("Pressure:", presHPa, "HPa")

	//Use datos de tipo string porque los datos sin las unidades no estarian completos
}

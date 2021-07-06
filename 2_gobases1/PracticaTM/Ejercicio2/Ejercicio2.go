package main

import "fmt"

func main() {
	var moisure float64
	var temperature float64
	var pressure float64

	moisure = 56
	temperature = 18
	pressure = 1016

	fmt.Println("La temperatura en Fco Álvarez es:", temperature, "ºC")
	fmt.Println("La humedad en Fco Álvarez es:", moisure, "%")
	fmt.Println("La presion en Fco Álvarez es:", pressure, "mb")
}

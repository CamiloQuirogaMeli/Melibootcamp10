package main

import "fmt"

func main() {
	var temperature float32
	var pressure float32
	var humidity int

	temperature = 10.5
	pressure = 1050.29
	humidity = 59

	fmt.Println("Temperatura: ", temperature, "C°")
	fmt.Println("Presion: ", pressure, "kPa")
	fmt.Println("Humedad: ", humidity, "%")
}
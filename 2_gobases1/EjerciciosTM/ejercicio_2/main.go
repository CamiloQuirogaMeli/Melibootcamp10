package main

import (
	"fmt"
)

func main() {
	var temperatura float32 = 15.3
	var humedad int = 81
	var presion float32 = 1007.62

	fmt.Printf("Temperatura: %g ºC\n", temperatura)
	fmt.Printf("Humedad: %d \n", humedad)
	fmt.Printf("Presion: %g hPa", presion)
}

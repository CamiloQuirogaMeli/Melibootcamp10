package main

import "fmt"

func main() {
	var (
		temp    = 17
		humedad = 47
		presion = 1016
	)
	fmt.Println(fmt.Sprintf("Temperatura: %v C", temp))
	fmt.Println(fmt.Sprintf("Humedad: %v %%", humedad))
	fmt.Println(fmt.Sprintf("Presion: %v hPa", presion))
}

package main

import "fmt"

func main() {

	var (
		temperatura float32
		humedad     float32
		presion     float32
	)

	temperatura = 25.4
	humedad = 45.4
	presion = 300.4

	fmt.Println("temperatura:", temperatura)
	fmt.Println("humedad:", humedad)
	fmt.Println("presion:", presion)
}

package main

import (
	"fmt"
)

func main() {
	var humedad, presion int = 50, 1080
	var temperatura float64 = 38.2
	fmt.Printf("La temperatura es %f, la presion es de %d, la humedad es de %d\n", temperatura, presion, humedad)
}

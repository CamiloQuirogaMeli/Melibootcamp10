package main

import (
	"fmt"
)

func main() {
	var (
		temperature, humidity int    = 30, 68
		pressure              string = "1012 MB"
	)
	fmt.Println("Temperatura ", temperature)
	fmt.Println("Humedad ", humidity)
	fmt.Println("Presion ", pressure)
}

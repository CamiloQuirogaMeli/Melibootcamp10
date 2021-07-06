package main

import (
	"fmt"
)

func main() {

	var temp int16 = 33
	var humidity int16 = 70
	var pressure float32 = 1020.3

	fmt.Printf("La temperatura en esta ciudad es %v C, la humedad es del %v y la presión %v   ", temp, humidity, pressure)

}

//en la docu dice que podemo poner v como grupo de tipos

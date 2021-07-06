package main

import (
	"fmt"
)

func main() {
	var temperatura float64 = 15.5 //Elija la temperatura
	var humedad float64 = 55       //Elija la humedad
	var presion int64 = 1022       // elija la presion
	fmt.Println("La temperatura hoy es de: ", temperatura)
	fmt.Println("La humedad de hoy es de: ", humedad, "%")
	fmt.Println("La presion de hoy es de: ", presion, "hPa")

}

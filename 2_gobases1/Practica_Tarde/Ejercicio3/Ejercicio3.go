package main

import "fmt"

func main() {
	var (
		edad       = 23
		empleado   = true
		antiguedad = 2
		sueldo     = 100000
	)

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Print("Cumple con los requisitos para el prestamo, sin interes!")
		} else {
			fmt.Print("Cumple con los requisitos para el prestamo, con interes!")
		}
	} else {
		fmt.Print("No cumple los requisitos para el prestamo")
	}

}

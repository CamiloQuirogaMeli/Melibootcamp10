package main

import "fmt"

func main() {
	var (
		edad       int     = 23
		empleado   bool    = true
		antiguedad int     = 2
		sueldo     float64 = 90000
	)

	switch {
	case edad < 22:
		fmt.Println("Lo sentimos, eres menor de 22 años")
	case empleado == false:
		fmt.Println("Lo sentimos, no eres empleado")
	case antiguedad < 1:
		fmt.Println("Lo sentimos, no tienes más de un año de antigüedad")
	case sueldo > 100000:
		fmt.Println("Felicitaciones! eres apto para otorgarte un préstamo y como tienes un sueldo mayor a $100.000, no te cobraremos intereses!")
	default:
		fmt.Println("Felicitaciones! eres apto para otorgarte un préstamo")
	}
}

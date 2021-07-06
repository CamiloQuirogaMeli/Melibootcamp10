package main

import "fmt"

func main() {
	var edad int = 24
	var trabaja bool = true
	var antiguedad int = 4
	var sueldo int = 80000
	if edad <= 22 || !trabaja || antiguedad <= 1 {
		fmt.Println("El cliente no cumple con los requisitos para recibir un prestamo.")
	} else if sueldo <= 100000 {
		fmt.Println("El cliente puede recibir un préstamo con intereses")
	} else {
		fmt.Println("El cliente puede recibir un préstamo sin intereses")
	}
}

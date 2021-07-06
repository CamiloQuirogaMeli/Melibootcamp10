package main

import "fmt"

func main() {
	var edad, antiguedad, sueldo int
	fmt.Println("Ingrese su edad")
	fmt.Scanln(&edad)
	fmt.Println("Ingrese su antiguedad laboral en años (en caso de no llegar al año trabajado ingrese 0)")
	fmt.Scanln(&antiguedad)
	fmt.Println("Ingrese su sueldo")
	fmt.Scanln(&sueldo)
	switch {
	case edad < 18:
		fmt.Println("Lo siento, debe ser mayor de edad para acceder a un prestamo")
	case antiguedad < 1:
		fmt.Println("Lo siento, debe tener al menos un año de antiguedad laboral para acceder a un prestamo")
	case sueldo > 100000:
		fmt.Println("Estimado cliente, usted cumple las condiciines para acceder a nuestros prestamos SIN intereses")
	default:
		fmt.Println("Estimado cliente, es apto para acceder a nuestros prestamos CON intereses")
	}
}

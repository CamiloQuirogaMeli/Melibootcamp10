package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float32 = 45857.90
	var nombre string = "Julián"

	fmt.Println("Nombre: " + nombre + "\nApellido: " + apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Valor de bool: ", boolean)
	fmt.Println("Sueldo: $", sueldo)
}

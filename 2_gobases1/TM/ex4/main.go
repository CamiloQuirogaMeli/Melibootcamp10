package main

import (
	"fmt"
)

func main() {
	var apellido string = "Gomez"
	var edad int = 35             // los int se declaran sin ""
	boolean := false              // los bool se declaran sin ""
	var sueldo float64 = 45857.90 // no usar string
	var nombre string = "Julián"

	// Para usar las variables y que compile
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(boolean)
	fmt.Println(sueldo)
	fmt.Println(nombre)
}

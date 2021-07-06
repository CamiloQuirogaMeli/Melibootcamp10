package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	// 1. Las comillas se usan con strings
	var edad int = 35
	// 2. En booleano tampoco se usan comillas
	boolean := false
	// 3. El tipo no coincide con la asignación. En este caso es un float
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}

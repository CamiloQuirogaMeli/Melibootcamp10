package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	//mal asignacion del tipo entero "35"
	var edad int = 35
	boolean := false
	//mal tipo de dato, debia ser un float
	var sueldo float64 = 458579.95
	var nombre string = "Julián"

	fmt.Println(apellido, edad, boolean, sueldo, nombre)

}

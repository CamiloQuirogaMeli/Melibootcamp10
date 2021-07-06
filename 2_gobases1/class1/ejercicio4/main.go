package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	// var edad int = "35" Cambiar "35" que es un string por 35 en modo numero
	var edad int = 35
	// boolean := "false"; aca al estar declarando una variable en vuelo no es necesario poner el tipo de dato y al ser un booleano hay que ponerlo sin comillas y sin punto y coma
	dato := false
	// var sueldo string = 45857.90 aqui el error es que se esta indicado que el tipo de dato era string y se inicializa con un float
	var sueldo float64 = 45857.90
	// correcta
	var nombre string = "Julian"

	fmt.Println(apellido, edad, dato, sueldo, nombre)
}

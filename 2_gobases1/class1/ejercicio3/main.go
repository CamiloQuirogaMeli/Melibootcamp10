package main

import "fmt"

func main() {

	// var 1nombre string Es incorrecta, seria:
	var nombre string
	// correcta
	var apellido string
	// var int edad Es incorrecta, seria:
	var edad int
	// 1apellido := 6 Es incorrecta, seria:
	//apellido := ""
	// var licencia_de_conducir = true Es incorrecta, seria:
	var licenciaDeConducir bool = true
	// var estatura de la persona int Es incorrecta, seria:
	var estaturaDeLaPersona int
	// correcta
	cantidadDeHijos := 2

	fmt.Println(nombre, apellido, edad, licenciaDeConducir, estaturaDeLaPersona, cantidadDeHijos)
}

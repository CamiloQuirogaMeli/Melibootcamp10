package main

import "fmt"

func ej4() {
	var apellido string = "Gomez" //está correcta

	//var edad int = "35" //el tipo de dato in no lleva ""
	//corrección
	var edad int = 35

	boolean := "false" //si el objetivo es declarar una variable boolean, primero no debe poner como nombre de
	//variable un nombre restringido como boolean, y aparte no van las "" ni el ;
	//corrección
	varBoolean := false

	//var sueldo string = 45857.90 //una variable sueldo debe ir en float
	//corrección
	var sueldo float32 = 45857.90

	var nombre string = "Julián" //está correcta

	fmt.Println(apellido, edad, boolean, varBoolean, sueldo, nombre)
}

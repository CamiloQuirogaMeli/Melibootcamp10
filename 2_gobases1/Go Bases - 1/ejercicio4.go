package main

import "fmt"

/*
	Ejercicio 4 - Tipos de datos
	Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos
	tipos en Go pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y
	pidió la ayuda de un desarrollador experimentado que pueda:
	Verificar su código y realizar las correcciones necesarias.
		var apellido string = "Gomez"
  		var edad int = "35"
  		boolean := "false";
  		var sueldo string = 45857.90
  		var nombre string = "Julián"
*/

func main() {
	//var apellido string = "Gomez" (CORRECTO)
	var apellido string = "Gomez"
	//var edad int = "35" (INCORRECTO)
	var edad int = 35
	//boolean := "false"; (INCORRECTO)
	boolean := false
	//var sueldo string = 45857.90 (INCORRECTO)
	var sueldo string = "45857.90"
	//var nombre string = "Julián" (CORRECTO)
	var nombre string = "Julián"
	fmt.Println(nombre, apellido, edad, sueldo, boolean)
}

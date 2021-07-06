package main

import (
	"fmt"
)

func main() {

	/*
		Un estudiante de programación intentó realizar declaraciones de variables con sus
		respectivos tipos en Go pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó
		su código y pidió la ayuda de un desarrollador experimentado que pueda:
			Verificar su código y realizar las correcciones necesarias.
	*/

	/*
		var apellido string = "Gomez"
		var edad int = "35"
		boolean := "false"
		var sueldo string = 45857.90
		var nombre string = "Julián"
	*/

	var apellido string = "Gomez"
	//var edad int = "35"
	boolean := "false"
	//var sueldo string = 45857.90
	var nombre string = "Julián"

	fmt.Println("Las variables apellido, boolean y nombre no tiene errores de sintaxis")
	fmt.Println("Ejecución: ", apellido, boolean, nombre)
	fmt.Println("Sin embargo en el caso de boolean, se declaró como un string, se sugiere declaralo como un valor booleano eliminando las comillas")

	fmt.Println("Para la variable edad se sugiere eliminar las comillas para que sea un entero")
	fmt.Println("Para la variable sueldo, se sugiere cambiar el tipo de variable de string a float16")

}

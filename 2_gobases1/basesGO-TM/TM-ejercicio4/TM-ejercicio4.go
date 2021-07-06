package main

import "fmt"

/*

VARIABLE 1
var apellido string = "Gomez"
CORRECTO

VARIABLE 2
var edad int = "35"
INCORRECTO -> El tipo de dato int no lleva comillas en su asignación
Lo correcto sería: var edad int = 35

VARIABLE 3
boolean := "false";
INCORRECTO -> El lenguaje GO no finaliza las lineas con punto y coma ;
La asignación correcta es: boolean := "false"
Estamos definiendo un string con el contenido "false"

VARIABLE 4
var sueldo string = 45857.90
INCORRECTO -> El tipo de dato string lleva comillas en su asignación ""
La asignación correcta es: var sueldo string = "45857.90"

VARIABLE 5
var nombre string = "Julián"
CORRECTO

*/

func main() {

	/*VARIABLE 1*/
	var apellido string = "Gomez"
	fmt.Println("Apellido: ", apellido)

	/*VARIABLE 2*/
	var edad int = 35
	fmt.Println("Edad: ", edad)

	/*VARIABLE 3*/
	boolean := "false"
	fmt.Println("Palabra: ", boolean)

	/*VARIABLE 4*/
	var sueldo string = "45857.90"
	fmt.Println("Sueldo: ", sueldo)

	/*VARIABLE 5*/
	var nombre string = "Julián"
	fmt.Println("Nombre: ", nombre)

}

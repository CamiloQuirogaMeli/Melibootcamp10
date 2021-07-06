package main

import "fmt"

/*

VARIABLE 1
var 1nombre string
INCORRECTO -> El nombre de una variable no puede comenzar con un numero
Una forma correcta de definir una variable nombre, es: var nombre string

VARIABLE 2
var apellido string
CORRECTO

VARIABLE 3
var int edad
INCORRECTO -> La variable se declara como: palabra reservada + nombre variable + tipo
Lo correcto seria: var edad int

VARIABLE 4
1apellido := 6
INCORRECTO -> El nombre de una variable no puede comenzar con un numero
Una forma válida de declarar esto es: apellido := 6, cabe destacar que
estamos definiendo un entero apellido y no un string, no tiene demasiado sentido logico, aún así es posible hacerlo

VARIABLE 5
var licencia_de_conducir = true
CORRECTO

VARIABLE 6
var estatura de la persona int
INCORRECTO -> El nombre de una variable no puede contener espacios
Una forma correcta de definir esto puede ser: var estatura_de_la_persona int
Aquí nuevamente hay otra cuestión de sentido lógico, la estatura la medimos con decimales, por lo tanto
lo ideal sería un valor flotante.

VARIABLE 7
cantidadDeHijos := 2
CORRECTO

*/

func main() {

	/*VARIABLE 1*/
	var nombre string
	nombre = "Agostina"
	fmt.Println("Nombre: ", nombre)

	/*VARIABLE 2*/
	var apellido string
	apellido = "Reinhardt"
	fmt.Println("Apellido: ", apellido)

	/*VARIABLE 3*/
	var edad int
	edad = 28
	fmt.Println("Edad: ", edad)

	/*VARIABLE 4*/
	apellido2 := 6
	fmt.Println("Dato: ", apellido2)

	/*VARIABLE 5*/
	var licencia_de_conducir = true
	fmt.Println("Licencia de conducir: ", licencia_de_conducir)

	/*VARIABLE 6*/
	var estatura_de_la_persona float32
	estatura_de_la_persona = 1.70
	fmt.Println("Estatura de la persona: ", estatura_de_la_persona)

	/*VARIABLE 7*/
	cantDeHijos := 2
	fmt.Println("Cantidad de hijos: ", cantDeHijos)

}

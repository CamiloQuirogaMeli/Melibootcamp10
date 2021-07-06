package main

import "fmt"

/*
	Ejercicio 3 - Declaración de variables
	Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia
	Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos
	del examen consiste en declarar distintas variables.
	Necesita ayuda para:
	Detectar cuáles de estas variables que declaró el alumno son correctas.
	Corregir las incorrectas.
		var 1nombre string
  		var apellido string
  		var int edad
  		1apellido := 6
  		var licencia_de_conducir = true
  		var estatura de la persona int
  		cantidadDeHijos := 2
*/

func main() {
	//var 1nombre string (INCORRECTO)
	var nombre string
	nombre = "Anderson Yessid"
	//var apellido string (CORRECTO)
	var apellido string
	apellido = "Torres"
	//var int edad (INCORRECTO)
	var edad int
	edad = 24
	//1apellido := 6 (INCORRECTO)
	_apellido := "Quijano"
	//var licencia_de_conducir = true (INCORRECTO)
	var licenciaDeConducir bool
	licenciaDeConducir = true
	//var estatura de la persona int (INCORRECTO)
	var estaturaDeLaPersona float64
	estaturaDeLaPersona = 1.75
	//cantidadDeHijos := 2 (CORRECTO)
	cantidadDeHijos := 2

	fmt.Println("El señor", nombre, apellido, _apellido, "tiene", edad, "años.")
	fmt.Println("¿Cuenta con licencia para conducir?:", licenciaDeConducir)
	fmt.Println("¿La estatura es?:", estaturaDeLaPersona)
	fmt.Println("¿Cuantos hijos tiene?:", cantidadDeHijos)
}

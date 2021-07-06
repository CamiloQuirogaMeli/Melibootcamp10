package main

import "fmt"

func main() {
	//printLnameAndAdress()
	weather()
}

func printLnameAndAdress() {
	var name string = "Martin Marcante"
	var address string = "Crespo 231, Sunchales, Santa Fe, Argentina"

	fmt.Println(name + " live in " + address)
}

func weather() {
	var temperature float64 = 20.3
	var humidity float64 = 60
	var pressure float64 = 1010

	fmt.Println("Weather in Sunchales: temperature: ", temperature,
		"Humidity: ", humidity, "%",
		"Pressure: ", pressure)
}

//DECLARACIÓN DE VARIABLES

//var 1nombre string INCORRECTA
var nombre string //sin numeros al comienzo

//var apellido string CORRECTA

//var int edad INCORRECTA
var edad int // este es el orden correcto

//1apellido := 6 INCORRECTA
apellido := 6 //sin numeros al comienzo (además de tener un error semántico al asignar un número a un supuesto apellido)

//var licencia_de_conducir = true CORRECTA

//var estatura de la persona int INCORRECTA
var estaturaDeLaPersona int //no se permiten espacios

// cantidadDeHijos := 2 CORRECTA
//---------------------------------------------------------------------

//TIPOS DE DATOS
//var apellido string = "Gomez" CORRECTA

//var edad int = "35" INCORRECTA
var edad int = 35 //era incorrecto el tipo de dato

//boolean := "false" CORRECTA

//var sueldo string = 45857.90 INCORRECTA
var sueldo float64 = 45857.90 //era incorrecto el tipo de dato

//var nombre string = "Julian" CORRECTA
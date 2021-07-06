package main

import "fmt"

func main() {
	//Ejercicio 1
	var nombre string = "Martina"
	var direccion string = "Florencio Sanches 817"
	fmt.Println("Nombre: " + nombre + "\n" + "Direccion: " + direccion)

	//Ejercicio 2
	var temperatura float64 = 16
	var humedad uint = 48
	var presion float64 = 1016.6
	fmt.Println("Temperatura: ", temperatura, "\n", "Humedad: ", humedad, "\n", "Presion: ", presion)
	//Ejercicio 3
	/*
		var 1nombre string
		var apellido string
		var int edad
		1apellido := 6
		var licencia_de_conducir = true
		var estatura de la persona int
		cantidadDeHijos := 2
	*/
	//Correccion
	/*
		var nombre string
		var apellido string //correcta
		var edad int
		apellido = "Apellido"
		var licenciaDeConducir bool = true
		var estarturaDeLaPersona int
		cantidadDeHijos := 2 //correcta
	*/

	//Ejercicio 4
	/*
		var apellido string = "Gomez" //correcta
		var edad int = 35
		bandera := false
		var sueldo string = "45857.90"
		//o
		var sueldo float64 = 45857.90
		var nombre string = "Julian" //correcta
	*/
}

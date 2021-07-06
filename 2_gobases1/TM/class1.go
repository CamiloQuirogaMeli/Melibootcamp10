package main

import "fmt"

func main() {
	//primer ejercicio
	/*
		nombre, direccion := "Emilia", "Las Heras - Mendoza"

		fmt.Println("Mi nombre es: " + nombre + " y vivo en " + direccion)
	*/

	//segundo ejercicio
	/*temperatura := 18.0
	humedad := 34.0
	presion := 1009.5

	fmt.Println(" La temperatura en Mendoza es: ", temperatura,
		" grados, mientras que la húmedad es de: ", humedad,
		" con una presión atmosférica de ", presion)
	*/

	//tercer ejercicio

	/*var lnombre string
	var apellido string

	//ERROR VA PRIMERO EL NOMBRE DE LA VARIABLE Y DESPUÉS EL TIPO
	//var int edad
	var edad int
	lapellido := 6
	apellido = "perez"
	var licencia_de_conducir = true

	//ERROR AL DECLARAR EL NOMBRE DE LA VARIABLE CON ESPACIOS
	//var estatura de la persona int
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	fmt.Println(lnombre+
		apellido,
		edad,
		lapellido,
		licencia_de_conducir,
		estatura_de_la_persona,
		cantidadDeHijos)
	*/

	//cuarto ejercicio

	var apellido string = "Gomez"

	//ERROR LA EDAD ESTÁ COMO INT Y HA INSTANCIADO UN STRING
	//var edad int = "35"
	var edad int = 35
	boolean := "false"

	//ERROR HA DECLARO EL SUELDO COMO UN STRING Y LO INSTANCIA COMO UN FLOAT
	//var sueldo string = 45857.90
	var sueldo string = "45857.90"
	var nombre string = "Julián"

	fmt.Println(apellido,
		edad,
		boolean,
		sueldo,
		nombre)

}

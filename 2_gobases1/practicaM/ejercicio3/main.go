package main

import "fmt"

func main() {
	//nombre de variables no inician con numeros
	var primerNombre string
	var apellido string
	//error orden de declaracion
	var edad int
	//nombre de variables no inician con numeros
	primerApellido := 6
	//falta tipo de dato
	var licenciaDeConducir bool = true
	//nombre de variables no tienen espacions
	var estaturaDeLaPersona int
	cantidadDeHijos := 2

	fmt.Println(primerNombre, apellido, edad, primerApellido, licenciaDeConducir, estaturaDeLaPersona, cantidadDeHijos)

}

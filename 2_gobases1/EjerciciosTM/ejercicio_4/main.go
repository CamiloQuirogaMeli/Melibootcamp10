package main

import (
	"fmt"
)

func main() {
	// declaracion y asiganacion de valor correctas
	var apellido string = "Gomez"
	fmt.Println("apellido: ", apellido)

	// aqui no hay error de sintaxis pero si es necesario decidir el tipo de dato
	// que se le quiere dar a la variable porque el valor que se le quiere asignar es incorrecto
	// segun el tipo definido
	// var edad int = "35";
	var edad int = 35
	fmt.Println("edad: ", edad)

	// aqui no hay error. Quizas confunda el nombre de la variable ya que es una palabra reservada
	// en algunos lenguajes de programacion. Cuesta dicernir entre que tipo y valor de variable de quiere
	// si un buleano o un string
	opcionBuleana := false
	opcionString := "false"
	fmt.Println("opcionBuleana: ", opcionBuleana)
	fmt.Println("opcionString: ", opcionString)

	// nuevamente el tipo de dato no coincide con el valor asignado
	// en este caso, segun el nombre de la variable, el tipo de dato debiese ser float32 o float64
	// var sueldo string = 45857.90
	var sueldo float64 = 45857.90
	fmt.Println("sueldo: ", sueldo)

	// definicion y asignacion correcta de la variable nombre
	var nombre string = "Julián"
	fmt.Println("nombre: ", nombre)
}

package main

import "fmt"

func ej3() {
	var lname string //está correcta

	var apellidos string //está bien pero repite el nombre de la variable anterior ya que lname es last name = apellido
	//corrección
	var fname string

	//var int edad //la declaración de variable está mal
	//corrección
	var edad int

	lapellido := 6 //por el nombre de la variable debería ser un string, pero si lo toma como int esá bien
	//corrección
	lapellidos := "6" //aún así no entiendo a que se refiere la variable

	var licencia_de_conducir = true //le falta poner el tipo de dato de la variable
	//corrección
	var licencia_de_conducir1 bool = true

	var estatura_de_la_persona int //el nombre de la variable debería ir todo junto, no puede ir separado
	//corrección
	var estatura_de_la_persona1 int

	cantidadDeHijos := 2 //está correcta

	fmt.Println(lname, apellidos, fname, edad, lapellido, lapellidos, licencia_de_conducir, licencia_de_conducir1)
	fmt.Println(estatura_de_la_persona, estatura_de_la_persona1, cantidadDeHijos)
}

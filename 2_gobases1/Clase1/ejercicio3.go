package main

import (
	"fmt"
)

func main() {

	/*
		Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia
		Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos
		del examen consiste en declarar distintas variables.
			Necesita ayuda para:
		1. Detectar cuáles de estas variables que declaró el alumno son correctas.
		2. Corregir las incorrectas.
	*/

	/*
		var 1nombre string
		var apellido string
		var int edad
		1apellido := 6
		var licencia_de_conducir = true
		var estatura de la persona int
		cantidadDeHijos := 2
	*/

	fmt.Println("Las variables correctas son las siguientes: ")

	fmt.Println("apellido")
	fmt.Println("cantidadDeHijos")
	fmt.Println()
	fmt.Println("Las variables incorrectas son las siguientes: ")

	fmt.Println("1nombre, porque las variables no pueden iniciar por números")
	fmt.Println("edad, porque el nombre de la variables es primero que su tipo")
	fmt.Println("1apellido, porque las variables no pueden iniciar por números")
	fmt.Println("licencia_de_conducir, porque para asignarle un valor, se utiliza :=")
	fmt.Println("estatura de la persona, porque las variables no llevan espacios")
	fmt.Println()

	fmt.Println("Una forma correcta de definirlas sería la siguiente")
	fmt.Println("var nombre1 string")
	fmt.Println("var edad int")
	fmt.Println("var apellido1 string")
	fmt.Println("licencia_de_conducir := true")
	fmt.Println("var estatura_de_la_persona int")
	fmt.Println()
}

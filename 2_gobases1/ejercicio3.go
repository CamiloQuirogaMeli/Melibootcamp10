package main

import "fmt"

//var 1nombre string una variable tiene que empezar con letras

var apellido string = "martinez"

//var int edad el tipo de variable tiene que ir en la posicion 2

//1apellido := 6 la variable no puede empezar con numeros

var licencia_de_conducir = true

//var estatura de la persona int el nombre de la variable no puede tener espacios

//cantidadDeHijos := 2 no se utilizo la palabra reservada var

func ejercicio3() {
	fmt.Println(apellido)
	if licencia_de_conducir {
		fmt.Println("Entro")
	}
}

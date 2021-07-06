package main

import (
	"fmt"
)

func main() {
	var lnombre string
	var apellido string
	var edad int // Antes: var int edad (orden incorrecto)
	lapellido := 6
	var licencia_de_conducir = true
	var estatura_de_la_persona int // Antes: var estatura de la persona int (las variables no deben tener espacios)
	cantidadDeHijos := 2

	// Para usar las variables y que compile
	fmt.Println(lnombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(lapellido)
	fmt.Println(licencia_de_conducir)
	fmt.Println(estatura_de_la_persona)
	fmt.Println(cantidadDeHijos)
}

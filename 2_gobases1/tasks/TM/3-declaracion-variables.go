package main

import "fmt"

func main() {
	var lnombre string
	var apellido string
	// 1. El tipado de una variable va después del nombre
	var edad int
	lapellido := 6
	// 2. Por buena práctica, uso de camelcase. Adicional, tipado de la variable
	var licenciaDeConduccion bool = true
	// 3. Sin espacios en la declaración de una variable
	var estaturaDeLaPersona int
	cantidadDeHijos := 2

	fmt.Println(
		lnombre,
		apellido,
		edad,
		lapellido,
		licenciaDeConduccion,
		estaturaDeLaPersona,
		cantidadDeHijos,
	)

}

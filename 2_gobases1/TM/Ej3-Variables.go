package main

import "fmt"

func main() {
	var nombre string = "Ignacio"
	var apellido string
	var edad int
	edad = 24
	apellido = "Gonzalez"
	var licenciaConducir bool = true
	var estatura float32 = 181.0
	cantidadDeHijos := 0

	fmt.Println("Nombre: " + nombre + "\nApellido: " + apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Licencia de conducir: ", licenciaConducir)
	fmt.Println("Estatura: ", estatura)
	fmt.Println("Cantidad de hijos: ", cantidadDeHijos)
}

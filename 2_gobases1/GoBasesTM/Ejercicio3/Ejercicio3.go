package main

import "fmt"

func main() {
	//var 1nombre string - No puede arrancar con numero
	var apellido string = "Esposito"
	//var int edad - Esta al reves
	//1apellido := 6 - No puede arrancar con nro
	var licencia_de_conducir = true
	//var estatura de la persona int no puede haber espacios
	cantidadDeHijos := 2

	fmt.Println("Apellido: ", apellido, " licencia_de_conducir ", licencia_de_conducir, "cantidadDeHijos ", cantidadDeHijos)
}

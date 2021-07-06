package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func main() {
	a := Alumnos{"Jonathan David", "Krucheski", "37950264", "22/06/21"}
	a.detalle()
}

func (a Alumnos) detalle() {
	fmt.Println("Nombre: ", a.Nombre)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
}

package main

import "fmt"

type Alumno struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

func (student Alumno) detalle() {
	fmt.Println("Nombre:", student.Nombre)
	fmt.Println("Apellido:", student.Apellido)
	fmt.Println("DNI:", student.DNI)
	fmt.Println("Fecha:", student.Fecha)
}

func main() {

	var student Alumno
	student.Nombre = "Homero"
	student.Apellido = "Simpson"
	student.DNI = 12345678
	student.Fecha = "11/11/11"

	student.detalle()
}
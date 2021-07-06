package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func main() {
	alumno := Alumno{Nombre: "Alex", Apellido: "Castillo", DNI: "1234567", Fecha: "29-06-2021"}
	alumno.detalle()
}

func (a Alumno) detalle() {
	fmt.Println(a)
}

package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) Detalle() {

	fmt.Printf("Alumno: %s %s\nDNI: %s\nFecha de ingreso: %s", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {

	alumno := Alumno{
		Nombre:   "Gianluca",
		Apellido: "Ciccarelli",
		DNI:      "12345678",
		Fecha:    "22/06/2021",
	}

	alumno.Detalle()
}

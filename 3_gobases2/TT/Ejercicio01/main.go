package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) detalle() {
	fmt.Println(a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alumno1 := Alumnos{"Ariel", "Gugliotta", 1234567, "23/05/1974"}
	alumno1.detalle()
}

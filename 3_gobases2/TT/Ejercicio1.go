package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string `json:"primer_nombre"`
	Apellido string `json:"apellido"`
	DNI      string `json:"dni"`
	Fecha    string `json:"fecha" bd:"fecha_inicio"`


}

func main() {
	student := Alumno{
		Nombre: "Conrado",
		Apellido: "Navarro",
		DNI: "40315",
		Fecha: "23/32/3221",
	}
	student.detalle()
}


func (std Alumno) detalle() {
	fmt.Println("Nombre: ", std.Nombre)
	fmt.Println("Apellido: ", std.Apellido)
	fmt.Println("DNI: ", std.DNI)
	fmt.Println("Fecha: ", std.Fecha)
}
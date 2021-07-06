package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (v Alumno) detalle() string {
	details := ""
	details += fmt.Sprintf("Nombre: \t[%s]\nApellido: \t[%s]\nDni: \t\t[%s]\nFecha: \t\t[%s]\n", v.Nombre, v.Apellido, v.Dni, v.Fecha)
	return details
}

func main() {
	student := Alumno{"Nicolas", "Barrera", "39.091.141", "23/05/1995"}
	fmt.Println(student.detalle())
}

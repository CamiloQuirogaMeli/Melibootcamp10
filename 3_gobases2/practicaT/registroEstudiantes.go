package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (e *Estudiante) setEstudiante(nombre string, apellido string, dni string, fecha string) {
	e.Nombre = nombre
	e.Apellido = apellido
	e.DNI = dni
	e.Fecha = fecha
}

func (e Estudiante) detalle() Estudiante {
	return e
}

func main() {
	e1 := Estudiante{"Laura", "Camelo", "1234567", "05-Julio-2020"}
	fmt.Println(e1.detalle())

	e1.setEstudiante("Vanessa", "Rodriguez", "98765432", "01-Enero-2015")
	fmt.Println(e1.detalle())

}

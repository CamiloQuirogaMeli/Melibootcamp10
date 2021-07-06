package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (a Alumnos) details() {
	fmt.Printf("Nombre: %s Apellido: %s DNI: %s Fecha: %s\n", a.Nombre, a.Apellido, a.Dni, a.Fecha)
}

func main() {
	a1 := Alumnos{"Carlos", "Molina", "1234567890", "2021-06-21"}
	a2 := Alumnos{"Juan", "Gutierrez", "1234567890", "2021-06-21"}
	a3 := Alumnos{"Juan Carlos", "Molina Gutierrez", "1234567890", "2021-06-21"}

	a1.details()
	a2.details()
	a3.details()

}

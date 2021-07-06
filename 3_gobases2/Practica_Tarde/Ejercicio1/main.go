package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalles() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %v\nFecha %s", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	a1 := Alumno{"Valeria", "Macina", 92915760, "29/06/2021"}
	a1.detalles()
}

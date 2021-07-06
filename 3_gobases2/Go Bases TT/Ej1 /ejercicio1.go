package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) Detalle() {
	fmt.Printf("Nombre[%v]\n", a.Nombre)
	fmt.Printf("Apellido[%v]\n", a.Apellido)
	fmt.Printf("Dni[%v]\n", a.DNI)
	fmt.Printf("Fecha[%v]\n", a.Fecha)
}

func main(a Alumno) {
	alumno := Alumno{"Juan", "Luna", "123456789", "20/03/2021"}

	alumno.Detalle()
}

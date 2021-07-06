package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) Detalle() {
	fmt.Println("Nombre", a.Nombre)
	fmt.Println("Apellido", a.Apellido)
	fmt.Println("DNI", a.DNI)
	fmt.Println("Fecha", a.Fecha)
}

func main() {
	a := Alumno{"Daniel", "Sanchez", "37629474", "21/06/2021"}
	a.Detalle()

}

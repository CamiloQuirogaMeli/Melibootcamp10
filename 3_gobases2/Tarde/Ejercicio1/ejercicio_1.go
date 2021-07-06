package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %s\nFecha:%s\n", a.Nombre, a.Apellido, a.Dni, a.Fecha)
}

func main() {
	a := Alumno{"Roman", "Riquelme", "12121212", "12/12/12"}
	a.detalle()
}

package main

import (
	"fmt"
)

type Alumno struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (alumno *Alumno) detalle() {
	fmt.Println("Nombre:", alumno.nombre)
	fmt.Println("Apellido:", alumno.apellido)
	fmt.Println("DNI:", alumno.dni)
	fmt.Println("Fecha:", alumno.fecha)
	fecha
}

func main() {
	alumno := Alumno{"Luis", "Oropeza", "iibj950507hmi", "95/05/07"}
	alumno.detalle()
}

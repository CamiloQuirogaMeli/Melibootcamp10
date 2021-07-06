package main

import (
	"fmt"
	"time"
)

type Alumnos struct {
	nombre   string
	apellido string
	dni      int
	fecha    time.Time
}

func (a Alumnos) detalle() {
	fmt.Println("Nombre ", a.nombre)
	fmt.Println("Apellido ", a.apellido)
	fmt.Println("DNI ", a.dni)
	fmt.Println("Fecha ", a.fecha)
}

func main() {
	a1 := Alumnos{
		nombre:   "Franco",
		apellido: "Ballet",
		dni:      51246045,
		fecha:    time.Now(),
	}
	a1.detalle()
}

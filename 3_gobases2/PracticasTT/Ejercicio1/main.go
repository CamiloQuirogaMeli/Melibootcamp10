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
	estudiante := Alumnos{
		nombre:   "Camilo",
		apellido: "Quiroga",
		dni:      123350302,
		fecha:    time.Now(),
	}
	estudiante.detalle()
}

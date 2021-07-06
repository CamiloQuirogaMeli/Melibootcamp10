package main

import (
	"fmt"
)

type Alumnos struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumnos) detalle() {

	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %v\nFecha: %s\n", a.nombre, a.apellido, a.dni, a.fecha)
}

func ejercicio6() {
	alumno := Alumnos{
		nombre:   "Federico",
		apellido: "Martinez",
		dni:      38168886,
		fecha:    "29 de junio",
	}

	alumno.detalle()
}

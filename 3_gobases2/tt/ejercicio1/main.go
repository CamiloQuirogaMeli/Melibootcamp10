package main

import "fmt"

type alumno struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (a alumno) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %s\nfecha: %s\n", a.nombre, a.apellido, a.dni, a.fecha)
}

func main() {
	alumno := alumno{
		nombre:   "Marcos",
		apellido: "Zabala",
		dni:      "40128284",
		fecha:    "20/01/2015",
	}
	alumno.detalle()
}

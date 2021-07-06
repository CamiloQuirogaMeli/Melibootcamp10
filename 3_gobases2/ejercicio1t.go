package main

import "fmt"

// Una universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el
// detalle de la siguiente manera:
// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]
// reemplazando los valores que están en corchetes por el valor de los alumnos.
// Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido,
// DNI, Fecha y que tenga un método detalle

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre: ", a.Nombre)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
}

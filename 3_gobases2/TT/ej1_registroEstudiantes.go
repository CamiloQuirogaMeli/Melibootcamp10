package main

import "fmt"

type Estudiante struct {
	Nombre       string
	Apellido     string
	DNI          int
	fechaIngreso string
}

func (e Estudiante) detalle() {
	fmt.Println(e.Nombre)
	fmt.Println(e.Apellido)
	fmt.Println(e.DNI)
	fmt.Println(e.fechaIngreso)
	fmt.Println("El alumno: ", e.Nombre, " ", e.Apellido, " ha sido ingresado exitosamente.")
}

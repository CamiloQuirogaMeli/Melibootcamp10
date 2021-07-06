package main

import (
	"fmt"
)

func main() {
	/*
		Ejercicio 1 - Impuestos de salario

		Una universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el
		detalle de la siguiente manera:
		Nombre: [Nombre del alumno]
		Apellido: [Apellido del alumno]
		DNI: [DNI del alumno]
		Fecha: [Fecha ingreso alumno]
		reemplazando los valores que están en corchetes por el valor de los alumnos.
		Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido,
		DNI, Fecha y que tenga un método detalle
	*/
	var alumnos = []Alumno{}
	a1 := Alumno{"Martina", "Olivera", 11111111, "22-06-2021"}
	alumnos = append(alumnos, a1)

	for i, _ := range alumnos {
		alumnos[i].detalle()
	}
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println(
		"Nombre: ", a.Nombre, "\n",
		"Apellido: ", a.Apellido, "\n",
		"DNI: ", a.DNI, "\n",
		"Fecha: ", a.Fecha)
}

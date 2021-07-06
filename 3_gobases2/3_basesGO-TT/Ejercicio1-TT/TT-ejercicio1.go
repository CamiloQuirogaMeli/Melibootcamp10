package main

import (
	"fmt"
)

/*
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

/*Defino estructura alumno*/

type alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

/* Metodo para detallar alumno */
func (a alumno) detalle() {

	fmt.Println("Nombre: ", a.nombre)
	fmt.Println("Apellido: ", a.apellido)
	fmt.Println("DNI: ", a.dni)
	fmt.Println("Fecha de ingreso: ", a.fecha)

}

/*Funcion lista alumno*/
func listaAlumnos(alumnos []alumno) {

	for _, a := range alumnos {
		a.detalle()
	}

}

func main() {

	/*Genero un slice de alumnos*/
	var alumnos []alumno
	a1 := alumno{"Nombre1", "Apellido1", 38748536, "22-06-21"}
	a2 := alumno{"Nombre2", "Apellido2", 41406510, "22-06-21"}
	a3 := alumno{"Nombre3", "Apellido3", 21738621, "22-06-21"}
	a4 := alumno{"Nombre4", "Apellido4", 22345821, "22-06-21"}
	alumnos = append(alumnos, a1, a2, a3, a4)

	/*Los muestro con el formato solicitado*/

	fmt.Println("Listado de alumnos")

	listaAlumnos(alumnos)

}

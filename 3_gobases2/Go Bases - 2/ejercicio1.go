package main

import "fmt"

/*
	Ejercicio 1 - Registro de estudiantes
	Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
	imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

	Nombre: [Nombre del alumno]
	Apellido: [Apellido del alumno]
	DNI: [DNI del alumno]
	Fecha: [Fecha ingreso alumno]

	Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
	alumnos/as.
	Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido,
	DNI, Fecha y que tenga un método detalle
*/

type Alumnos struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func detalles(a Alumnos) {
	fmt.Printf("Nombre:%s \nApellido:%s \nDNI:%d \nFecha: %s \n", a.nombre, a.apellido, a.dni, a.fecha)
}

func main() {
	a1 := Alumnos{
		nombre:   "Anderson",
		apellido: "Torres",
		dni:      1070,
		fecha:    "29-04-2021",
	}
	detalles(a1)
}

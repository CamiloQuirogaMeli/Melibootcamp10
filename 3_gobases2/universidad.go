package main

import "fmt"

//Una universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el detalle de la siguiente manera:
//Nombre: [Nombre del alumno]
//Apellido: [Apellido del alumno]
//DNI: [DNI del alumno]
//Fecha: [Fecha ingreso alumno]
//reemplazando los valores que están en corchetes por el valor de los alumnos.
//Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Alumno struct {
	Nombre string
	Apellido string
	DNI string
	Fecha string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre: [", a.Nombre, "]\nApellido: [" , a.Apellido, "]\nDNI: [", a.DNI, "]\nFecha: [", a.Fecha, "]")
}

func main()  {
	a1 := Alumno {
		Nombre: "Juanita",
		Apellido: "Benitez",
		DNI: "506171",
		Fecha: "15-jun-21",
	}

	a2 := Alumno {
		Nombre: "Paula",
		Apellido: "Ballen",
		DNI: "506170",
		Fecha: "15-jun-21",
	}

	a1.detalle()
	a2.detalle()
}
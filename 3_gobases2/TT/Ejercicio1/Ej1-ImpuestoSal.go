package main

import "fmt"

type Persona struct {
	Nombre string
	Apellido string
	DNI string
	Fecha string
}

func (p Persona) detalle(){
	fmt.Println(p.Nombre, p.Apellido, p.DNI, p.Fecha)
}

func main(){
	persona := Persona{"Ignacio","Gonzalez","12345678","16/10/96"}
	persona.detalle()
}

/*Una universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el
detalle de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
reemplazando los valores que están en corchetes por el valor de los alumnos.
Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido,
DNI, Fecha y que tenga un método detalle*/
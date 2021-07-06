package main

import (
	"fmt"
	"time"
)

/* Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle */

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          int64
	FechaIngreso time.Time
}

func (a *Alumno) detalle() {
	fmt.Printf("el metodo detalle se ejecuto \n")

}

func main() {

	var firstName string
	var lastName string
	var dni int64
	date := time.Now()

	fmt.Print("Ingrese el nombre del alumno: \n")
	fmt.Scanf("%v", &firstName)

	fmt.Print("Ingrese el apellido del alumno: \n")
	fmt.Scanf("%v", &lastName)

	fmt.Print("Ingrese el DNI del alumno: \n")
	fmt.Scanf("%v", &dni)

	a1 := Alumno{firstName, lastName, dni, date}

	fmt.Printf(" El alumno se creo con los siguientes datos: \n %v", a1)

	a1.detalle()

}

package main

import (
	"fmt"
	"reflect"
	"time"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func main() {
	/*
			universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el
		detalle de la siguiente manera:
		Nombre: [Nombre del alumno]
		Apellido: [Apellido del alumno]
		DNI: [DNI del alumno]
		Fecha: [Fecha ingreso alumno]
		reemplazando los valores que están en corchetes por el valor de los alumnos.
		Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido,
		DNI, Fecha y que tenga un método detalle
	*/
	alumno1 := Alumno{}
	fmt.Println("Digita los siguientes datos del alumno:")
	fmt.Print("Nombre: ")
	fmt.Scanln(&alumno1.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scanln(&alumno1.Apellido)
	fmt.Print("DNI: ")
	fmt.Scanln(&alumno1.DNI)
	alumno1.Fecha = time.Now()

	fmt.Println("\nLos datos del alumno son: \n")

	valores := reflect.ValueOf(alumno1)
	tipos := valores.Type()
	for i := 0; i < valores.NumField(); i++ {
		fmt.Println(tipos.Field(i).Name, ":", valores.Field(i))
	}

}

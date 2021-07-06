package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := Alumno{
		Nombre:   "Luciano",
		Apellido: "Dominino",
		DNI:      "38136456",
		Fecha:    "22-06-2021",
	}
	nombre, apellido, dni, fecha := a1.detail()

	a := reflect.TypeOf(a1)

	for i := 0; i < a.NumField(); i++ {
		var valor string
		field := a.Field(i)

		switch i {
		case 0:
			valor = nombre
		case 1:
			valor = apellido
		case 2:
			valor = dni
		case 3:
			valor = fecha
		}
		if valor != apellido {
			fmt.Printf("%v:\t\t %v\n", field.Name, valor)
		} else {
			fmt.Printf("%v:\t %v\n", field.Name, valor)
		}
	}
}

type Alumno struct {
	Nombre   string `bd:"primer_nombre"`
	Apellido string `bd:"apellido"`
	DNI      string `bd:"dni"`
	Fecha    string `bd:"fecha_ingreso"`
}

func (a Alumno) detail() (nombre, apellido, dni, fecha string) {
	return a.Nombre, a.Apellido, a.DNI, a.Fecha
}

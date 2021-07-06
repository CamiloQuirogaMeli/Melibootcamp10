package main

import (
	"fmt"
)

type Estudiante struct {
	nombre   string
	apellido string
	DNI      int
	fecha    Fecha
}

func (e Estudiante) detalle() {
	fmt.Println("Nombre: ", e.nombre)
	fmt.Println("Apellido: ", e.apellido)
	fmt.Println("DNI: ", e.DNI)
	fmt.Println("Fecha: ", e.fecha.day, "/", e.fecha.month, "/", e.fecha.year)
}

type Fecha struct {
	day   int
	month int
	year  int
}

func main() {
	var estudiante Estudiante
	fmt.Print("Escriba el nombre del estudiante: ")
	fmt.Scanln(&estudiante.nombre)
	fmt.Print("Escriba el apellido del estudiante: ")
	fmt.Scanln(&estudiante.apellido)
	fmt.Print("Escriba el DNI del estudiante: ")
	fmt.Scanln(&estudiante.DNI)
	fmt.Print("Escriba el dia (DD) de la fecha de ingreso del estudiante: ")
	fmt.Scanln(&estudiante.fecha.day)
	fmt.Print("Escriba el mes (MM) de la fecha de ingreso del estudiante: ")
	fmt.Scanln(&estudiante.fecha.month)
	fmt.Print("Escriba el año (YYYY) de la fecha de ingreso del estudiante: ")
	fmt.Scanln(&estudiante.fecha.year)
	estudiante.detalle()
}

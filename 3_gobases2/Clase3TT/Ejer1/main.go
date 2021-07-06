package main

import "fmt"

func main() {
	student := Alumn{"AAA", "BBB", "dsda", 123}
	student.detalle()
}

type Alumn struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    int64
}

func (a Alumn) detalle() {
	fmt.Println("Nombre ", a.Nombre)
	fmt.Println("Apellido ", a.Apellido)
	fmt.Println("DNI ", a.DNI)
	fmt.Println("Fecha ", a.Fecha)
}

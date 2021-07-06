package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre: ", a.Nombre)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
}

func main() {
	a1 := Alumno{"Carla", "Gomez", 35258954, "16/8/1989"}
	a2 := Alumno{}
	a2.Nombre = "Martin"
	a2.Apellido = "Perez"
	a2.DNI = 34259650

	a1.detalle()
	a2.detalle()

}

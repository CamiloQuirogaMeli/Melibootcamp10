package main

import "fmt"

type Alumno struct {
	Nombre    string
	Apellidos string
	DNI       int
	Fecha     string
}

func (alumno *Alumno) detalle() {
	fmt.Println(alumno)
}

func main() {

	alumno := Alumno{"Ulises", "Schreiner", 12345678, "29/07/2021"}

	alumno.detalle()

}

package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func main() {
	alumno1 := Alumno{"Juan", "Perez", 42123123, time.Now()}
	alumno1.detalle()
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %s\n", a.Nombre)
	fmt.Printf("Apellido: %s\n", a.Apellido)
	fmt.Printf("DNI: %d\n", a.DNI)
	fmt.Printf("Fecha: %s\n", a.Fecha)
}

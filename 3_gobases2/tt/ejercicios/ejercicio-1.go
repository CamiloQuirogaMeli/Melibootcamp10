package ejercicios

import (
	"fmt"
	"time"
)

type alumno struct {
	nombre       string
	apellido     string
	dni          int
	fechaIngreso time.Time
}

func (a alumno) detalle() {
	fmt.Printf("Nombre: %s\n", a.nombre)
	fmt.Printf("Apellido: %s\n", a.apellido)
	fmt.Printf("DNI: %d\n", a.dni)
	fmt.Printf("Fecha: %s\n\n", a.fechaIngreso)
}

func PrimerEjercicio() {
	alumnos := []alumno{
		{"Katherine", "Moreno", 1, time.Now()},
		{"Kevin", "Patiño", 2, time.Now()},
		{"Juan", "Falla", 3, time.Now()},
		{"Stiven", "Cardenas", 1, time.Now()},
	}

	for _, alumno := range alumnos {
		alumno.detalle()
	}
}

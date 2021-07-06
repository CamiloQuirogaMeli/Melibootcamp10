package main

import "fmt"

//Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go. Para ello requieren una estructura Matrix que tenga los métodos:
//- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
//- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea entre filas)
//La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es cuadrática y el valor máximo.

type Matrix struct {
	Matrix string
	Alto float64
	Ancho float64
	Cuadratica bool
	ValMax float64
}

func (m Matrix) detalle() {
	fmt.Println("Nombre: [", a.Nombre, "]\nApellido: [" , a.Apellido, "]\nDNI: [", a.DNI, "]\nFecha: [", a.Fecha, "]")
}

func main()  {
	a1 := Alumno {
		Nombre: "Juanita",
		Apellido: "Benitez",
		DNI: "506171",
		Fecha: "15-jun-21",
	}

	a2 := Alumno {
		Nombre: "Paula",
		Apellido: "Ballen",
		DNI: "506170",
		Fecha: "15-jun-21",
	}

	a1.detalle()
	a2.detalle()
}
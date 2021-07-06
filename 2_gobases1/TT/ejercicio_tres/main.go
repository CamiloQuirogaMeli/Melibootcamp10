package main

import (
	"fmt"
)

func main() {

	var (
		edad              int
		mesesDeAntiguedad int
		sueldo            int
		estaEmpleado      string
	)

	fmt.Print("Se encuentra empleado? Y/N: ")
	estaEmpleado = "N"
	fmt.Scanf("%s", &estaEmpleado)

	fmt.Print("Ingrese la edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Print("Meses de antiguedad: ")
	fmt.Scanf("%d", &mesesDeAntiguedad)

	fmt.Print("sueldo: ")
	fmt.Scanf("%d", &sueldo)

	switch {

	case (estaEmpleado != "N" && estaEmpleado != "Y"):
		fmt.Println("Opcion incorrecta, se tomara como no empleado por defecto")
		fallthrough
	case estaEmpleado == "N":
		fmt.Printf("El cliente no puede disponer de un prestamo")
	case (edad > 22 && mesesDeAntiguedad > 12 && sueldo > 100000):
		fmt.Printf("El cliente puede disponer de un prestamo sin intereses")
	case (edad > 22 && mesesDeAntiguedad > 12):
		fmt.Printf("El cliente puede disponer de un prestamo con intereses")
	default:
		fmt.Printf("El cliente no puede disponer de un prestamo")

	}

}

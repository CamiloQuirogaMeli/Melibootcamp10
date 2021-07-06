// Ejercicio 1 - Impuestos de salario

// Una universidad necesita registrar a los alumnos y generar una funcionalidad para imprimir el
// detalle de la siguiente manera:
// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]
// reemplazando los valores que están en corchetes por el valor de los alumnos.
// Para ello necesitamos generar una estructura Alumnos con las variables Nombre, Apellido,
// DNI, Fecha y que tenga un método detalle

package main

import "fmt"

func main() {

	type Alumno struct {
		Nombre   string
		Apellido string
		DNI      int
		Fecha    string
	}

	A1 := Alumno{"Martin", "Hernandez", 44642945, "04/04/86"}
	A2 := Alumno{"Juana", "Perez", 44645564, "06/07/89"}
	var Alumnos []Alumno
	Alumnos = append(Alumnos, A1)
	Alumnos = append(Alumnos, A2)
	fmt.Println()
	for i := 0; i < len(Alumnos); i++ {
		fmt.Printf("Nombre: %s\n", Alumnos[i].Nombre)
		fmt.Printf("Apellido: %s\n", Alumnos[i].Apellido)
		fmt.Printf("DNI: %d\n", Alumnos[i].DNI)
		fmt.Printf("Fecha: %s\n\n", Alumnos[i].Fecha)

	}

}

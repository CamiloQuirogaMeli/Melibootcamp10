/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido,
DNI, Fecha y que tenga un método detalle
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Persona struct {
	Nombre   string
	Apellido string
	Dni      string
	fecha    time.Time
	detalles Detalle
}

type Detalle struct{}

func main() {
	MENSAJE := `
	Bienvenido al programa de registro de estudiantes
	`
	fmt.Println(MENSAJE)
	fmt.Println("Antes de iniciar, ¿cuantos estudiantes desea registrar?")
	var numStudents int
	_, errNumStudents := fmt.Scanf("%d", &numStudents)

	if errNumStudents != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	students := []Persona{}
	for i := 1; i <= numStudents; i++ {
		students = append(students, Persona{})
		fmt.Println("DATOS DEL ESTUDIANTE Nº", i)
		fmt.Print("Nombre: ")
		scanner.Scan()
		students[i-1].Nombre = scanner.Text()
		fmt.Print("Apellido: ")
		scanner.Scan()
		students[i-1].Apellido = scanner.Text()
		fmt.Print("DNI: ")
		scanner.Scan()
		students[i-1].Dni = scanner.Text()
		fmt.Print("Fecha (yyyy-mm-dd): ")
		scanner.Scan()

		date, err := time.Parse("2006-01-02", scanner.Text())

		if err != nil {
			fmt.Println()
			fmt.Println("Formato de fecha invalida")
			os.Exit(0)
		}
		students[i-1].fecha = date
	}
	fmt.Println("Los datos de los estudiantes son:")
	for i, student := range students {
		fmt.Println("Los datos del estudiante", i, "son:")
		fmt.Print("Nombre: ", student.Nombre)
		fmt.Print(" Apellido: ", student.Apellido)
		fmt.Print(" DNI: ", student.Dni)
		fmt.Print(" Fecha: ", student.fecha)
	}
}

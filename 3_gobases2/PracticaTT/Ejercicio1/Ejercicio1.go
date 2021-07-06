// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

package main

import (
	f "fmt"
)

type student struct {
	name      string
	lastname  string
	dni       int
	dateEntry string
}

func detalle(students []student) {
	f.Printf("------------------------------------\n")
	for i, s := range students {
		f.Printf("Estudiante numero: %d\n", i+1)
		f.Printf("Nombre: %s\n", s.name)
		f.Printf("Apellido: %s\n", s.lastname)
		f.Printf("Dni: %d\n", s.dni)
		f.Printf("Fecha de ingreso: %s\n", s.dateEntry)
		f.Printf("------------------------------------\n")
	}
}

func main() {

	var cantStudents int
	var students []student
	var studentInfo student

	f.Printf("¿Cuantos estudiantes desea ingresar?: ")
	f.Scan(&cantStudents)

	for cantStudents <= 0 {
		f.Printf("Cantidad de estudiantes incorrecto, reingresar: ")
		f.Scan(&cantStudents)
	}

	for i := 0; i < cantStudents; i++ {

		f.Println("Información del estudiante", i+1)

		f.Printf("Ingrese el nombre del estudiante: ")
		f.Scan(&studentInfo.name)
		f.Printf("Ingrese el apellido del estudiante: ")
		f.Scan(&studentInfo.lastname)
		f.Printf("Ingrese el dni del estudiante: ")
		f.Scan(&studentInfo.dni)
		f.Printf("Ingrese la fecha de ingreso del estudiante (yyyy-mm-dd): ")
		f.Scan(&studentInfo.dateEntry)
		students = append(students, studentInfo)
	}
	detalle(students)
}

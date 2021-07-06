package main

import "fmt"

type student struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func main() {

	students := createMock()

	for index := range students {
		students[index].getDetail()
		fmt.Println()
	}

}

func (e student) getDetail() {
	fmt.Println("Nombre: ", e.nombre)
	fmt.Println("Apellido: ", e.apellido)
	fmt.Println("DNI: ", e.dni)
	fmt.Println("Fecha: ", e.fecha)
}

func createMock() (students []student) {
	students = append(students, student{"Javier", "Guggeri", 52563654, "22/06/2021"})
	students = append(students, student{"Martin", "Ayusto", 54567235, "05/01/2019"})
	return
}

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

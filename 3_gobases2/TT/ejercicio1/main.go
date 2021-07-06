package main

import (
	"fmt"
	"time"
)

type Student struct {
	name     string
	lastName string
	id       string
	date     time.Time
}

func (s Student) details() string {
	return "Nombre: " + s.name + "\nApellido: " + s.lastName + "\nDNI: " + s.id + "\nFecha: " + s.date.String()
}

func main() {
	var option int
	var aux int = 1
	const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

	const options string = "1.Buscar estudiante por nombre \n2. Agregar estudiante \n 0. Salir"

	students := make(map[string]Student)

	for aux != 0 {
		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var name string
			fmt.Println("Ingrese el nombre del estudiane")
			fmt.Scan(&name)

			student := students[name]
			fmt.Println("\n\n***** ESTUDIANTE **** \n" + student.details())

		case 2:
			var name, lastname, id string

			fmt.Println("Ingrese el nombre del estudiante")
			fmt.Scan(&name)
			fmt.Println("Ingrese el apellid del estudiante")
			fmt.Scan(&lastname)
			fmt.Println("Ingrese el DNI del estudiante")
			fmt.Scan(&id)

			students[name] = Student{name, lastname, id, time.Now()}

			fmt.Println("Usuario agregado con exito!")

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}

	}

}

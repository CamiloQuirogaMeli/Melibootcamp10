package main

import (
	"fmt"
)

var (
	students = []string{
		"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro",
		"Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}
)

func main() {

	fmt.Println("\n\tParte A")

	for i := range students {
		println(students[i])
	}

	fmt.Println("\n\tParte B")

	students = append(students, "Gabriela")

	for i := range students {
		println(students[i])
	}

}

/*
	a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
	necesario generar una aplicación que contenga dicha lista.
	Estudiantes:
	Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
	Leandro, Eduardo, Duvraschka.

	b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
	sin modificar el código que escribiste inicialmente.
	Estudiante:
	Gabriela
*/

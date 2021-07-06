// Ejercicio 5 - Listado de nombres

// a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
// necesario generar una aplicación que contenga dicha lista.
// Estudiantes:
// Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
// Leandro, Eduardo, Duvraschka.

// b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
// sin modificar el código que escribiste inicialmente.
// Estudiante:
// Gabriela

package main

import (
	"fmt"
	"strings"
)

func main() {

	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo",
		"Duvraschka"}
	var option string
	var newStudent string

	fmt.Println("La lista de estudiantes es la siguiente")

	for i := 0; i < len(students); i++ {
		fmt.Println("Estudiante:", students[i])
	}

	fmt.Println("¿Desea agregar un nuevo estudiante? (si/no) ")
	fmt.Scan(&option)
	option = strings.ToLower(option)

	for strings.Compare(option, "si") == 0 {
		fmt.Printf("ingrese el nombre del estudiante: ")
		fmt.Scan(&newStudent)
		students = append(students, newStudent)
		fmt.Println("¿Desea agregar un nuevo estudiante? (si/no) ")
		fmt.Scan(&option)
		option = strings.ToLower(option)
	}

	fmt.Println("La nueva lista de estudiantes es la siguiente")

	for i := 0; i < len(students); i++ {
		fmt.Println("Estudiante:", students[i])
	}

}

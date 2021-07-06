package main

import f "fmt"

func main() {
	/*
	 Ejercicio 5 - Listado de nombres

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

	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}
	f.Println("Estudiantes antes \n", students)
	students = append(students, "Gabriela")
	f.Println("Estudiantes despues \n", students)

}

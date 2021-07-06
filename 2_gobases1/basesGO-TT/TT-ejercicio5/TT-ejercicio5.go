package main

import "fmt"

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

func listaAlumnos(alumnos []string) {

	/*listo los alumnos de la clase que se encuentran en el slice*/

	fmt.Println("Alumnos de la clase: ")
	for i := 0; i < len(alumnos); i++ {

		fmt.Println(alumnos[i])

	}

}

func agregaAlumno(alumnos []string, nuevoAlumno string) {

	/*Agrego al alumno solicitado*/
	alumnos = append(alumnos, nuevoAlumno)
	listaAlumnos(alumnos)

}

func main() {

	alumnos := []string{"Benjamín", "Nahuel", "Brenda", "Marcos", "Pedro",
		"Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo",
		"Duvraschka"}

	listaAlumnos(alumnos)

	nuevoAlumno := "Gabriela"

	/*De este modo la estoy agregando temporal para listar*/
	agregaAlumno(alumnos, nuevoAlumno)

}

package ejercicios

import "fmt"

var estudiantes []string

func QuintoEjercicioLiteralA() {
	var primerosEstudiantes = []string{
		"Benjamin",
		"Nahuel",
		"Brenda",
		"Marcos",
		"Pedro",
		"Axel",
		"Alez",
		"Dolores",
		"Federico",
		"Hernan",
		"Leandro",
		"Eduardo",
		"Duvraschka",
	}
	estudiantes = primerosEstudiantes
	fmt.Println(estudiantes)
}

func QuintoEjercicioLiteralB() {
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println(estudiantes)
}

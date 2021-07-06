package main

import (
	"fmt"
)

func main() {

	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(estudiantes)

	// agregamos estudiante
	estudiantes = append(estudiantes, "Gabriela")

	fmt.Println(estudiantes)
}

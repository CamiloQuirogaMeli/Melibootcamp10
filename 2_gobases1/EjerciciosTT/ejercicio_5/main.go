package main

import (
	"fmt"
)

func main() {

	alumnos := []string{
		"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka",
	}

	fmt.Println("Alumnos: ", alumnos)

	alumnos = append(alumnos, "Gabriela")

	fmt.Println("Alumnos: ", alumnos)
}

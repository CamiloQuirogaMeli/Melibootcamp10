package main

import "fmt"

func listaEstudiantes() {

	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(append(estudiantes, "Gabriela"))

}

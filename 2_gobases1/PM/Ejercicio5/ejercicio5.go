package main

import "fmt"

func main() {

	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	for i, estudiante := range estudiantes {
		i = i
		fmt.Println(estudiante)
	}

	fmt.Println("--------------------------")

	estudiantes = append(estudiantes, "Gabriela")
	for i, estudiante := range estudiantes {
		i = i
		fmt.Println(estudiante)
	}
}

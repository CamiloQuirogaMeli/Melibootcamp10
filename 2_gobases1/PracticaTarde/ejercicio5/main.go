package main

import "fmt"

func main() {
	var newEstudiante string
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Ingrese nuevo estudiante:")
	fmt.Scanln(&newEstudiante)

	estudiantes = append(estudiantes, newEstudiante)

	fmt.Println(estudiantes)
}

package main

import "fmt"

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Lista de estudiantes:", estudiantes)
	fmt.Println("Luego de 2 clases...")
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println("Lista de estudiantes:", estudiantes)
}

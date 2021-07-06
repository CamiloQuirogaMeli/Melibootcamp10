package main

import "fmt"

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Los estudiantes inscriptos hasta ahora son:")
	fmt.Println(estudiantes)
	fmt.Println("Incorporando nuevo estudiante")
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println("Nuevo listado de estudiantes:")
	fmt.Println(estudiantes)
}

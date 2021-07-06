package main

import "fmt"

func main() {
	studentsList := []string{
		"Benjamin", "Nahuel", "Brenda",
		"Marcos", "Pedro", "Axel", "Alez",
		"Dolores", "Federico", "Hernan", "Leandro",
		"Eduardo", "Duvraschka"}

	fmt.Println("Lista de estudiantes:", studentsList)

	studentsList = append(studentsList, "Gabriela")
	fmt.Println("Lista de estudiantes con la persona nueva:", studentsList)
}

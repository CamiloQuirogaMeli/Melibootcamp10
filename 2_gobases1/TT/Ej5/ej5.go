package main

import "fmt"

func main() {
	// parte a
	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	// parte b
	students = append(students, "Gabriela")

	fmt.Println(students)
}

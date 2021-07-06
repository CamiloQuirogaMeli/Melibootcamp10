package main

import "fmt"

func main() {
	var students []string

	students = append(students, "Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores",
		"Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka")

	students = append(students, "Gabriela")

	fmt.Println(students)
}

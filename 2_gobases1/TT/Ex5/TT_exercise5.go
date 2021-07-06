package main

import "fmt"

func main() {

	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Students:", students)

	students = append(students, "Gabriela")

	fmt.Println("Students:", students)
}

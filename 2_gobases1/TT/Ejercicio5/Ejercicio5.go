package main

import "fmt"

func main() {

	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Estudiantes originales \n ", students)
	students = append(students, "Gabriela")
	fmt.Println("Estudiantes finales \n ", students)
}

package main

import "fmt"

func main() {
	students := []string{"Benjamin",
		"Nahuel",
		"Brenda",
		"Marcos",
		"Pedro",
		"Axel",
		"Alez",
		"Dolores",
		"Federico",
		"Hernan",
		"Leandro",
		"Eduardo",
		"Duvraschka"}
	students = append(students, "Gabriela")
	fmt.Println(students)
}

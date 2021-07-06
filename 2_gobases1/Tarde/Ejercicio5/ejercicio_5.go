package main

import "fmt"

func main() {

	//Item A
	var alumnos = []string{
		"Benjamin",
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

	//Item B
	alumnos = append(alumnos, "Gabriela")

	fmt.Println(alumnos)
}

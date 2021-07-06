package main

import "fmt"

func main() {
	findOrCreate("Gabriela")
}

func findOrCreate(student string) {

	var list = []string{
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
		"Duvraschka",
	}

	n := len(list)
	isOnList := false

	for i := 0; i < n; i++ {
		if student == list[i] {
			isOnList = true
		}
	}

	if !isOnList {
		list = append(list, student)
	}

	fmt.Println(list)
}

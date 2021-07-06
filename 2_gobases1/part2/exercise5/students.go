package main

import (
	f "fmt"
)

func main() {
	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	f.Println("Students: ")
	listNumber := 0;
	for _, student := range students {
		listNumber++
		f.Printf("%d.- %s\n",listNumber, student)
	}
	students = append(students, "Gabriela")
	f.Println("Students after 2 class: ")
	listNumber = 0;
	for _, student := range students {
		listNumber++
		f.Printf("%d.- %s\n",listNumber, student)
	}
}
package main

import "fmt"

var ()

func main() {

	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez",
		"Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	students = append(students, "Gabriela")

	printSlice(students)

}

func printSlice(s []string) {
	fmt.Printf("%v", s)
}

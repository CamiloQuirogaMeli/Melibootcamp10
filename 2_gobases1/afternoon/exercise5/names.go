package main

import "fmt"

func main() {

	initialList := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	newStudent := "Gabriela"
	fmt.Println("The current list of students is: ")

	printStudents(initialList)

	fmt.Println("Oh but wait, there is a new student called", newStudent, "\nAfter adding the new student, this is the final list:")

	finalList := append(initialList, newStudent)
	printStudents(finalList)

}

func printStudents(studentList []string) {
	for index, name := range studentList {
		fmt.Print(name)
		if index < len(studentList)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

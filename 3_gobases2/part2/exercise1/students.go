package main

import (
	f "fmt"
)

type Students struct {
	Name string
	LastName string
	DNI int
	Date string
}

func main() {
	var students []Students
	var name, lastName, date string
	var dni int
	var flag bool = true

	f.Println("Welcome, enter the info to register a student")

	for flag {
	f.Print("Name: ")
	f.Scanln(&name)
	f.Print("Lastname: ")
	f.Scanln(&lastName)
	f.Print("DNI: ")
	f.Scanln(&dni)
	f.Print("Date of admission (dd/mm/AAAA): ")
	f.Scanln(&date)

	student := Students{
		Name: name,
		LastName: lastName,
		DNI: dni,
		Date: date,
	}

	students = append(students, student)

	f.Print("Do you wamt another student? 1 = yes, 2 = no: ")
	f.Scanln(&flag)

	if !flag {
		break
	}
	}

	f.Println("List of students: ")
	for i , student := range students {
		f.Printf("Student %d.-\n", i + 1)
		f.Println("Name:", student.Name)
		f.Println("Lastname:", student.LastName)
		f.Println("DNI:", student.DNI)
		f.Println("Name:", student.Date)
		f.Println("**************************")
	}
}
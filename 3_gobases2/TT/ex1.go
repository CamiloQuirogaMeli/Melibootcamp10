package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name            string
	Surname         string
	DNI             string
	DateOfAdmission time.Time
}

func (student Student) details() {
	fmt.Println("Nombre:", student.Name)
	fmt.Println("Apellido:", student.Surname)
	fmt.Println("DNI:", student.DNI)
	fmt.Println("Fecha:", student.DateOfAdmission)
}

func main() {
	student := Student{
		Name:            "Carlos",
		Surname:         "Castro",
		DNI:             "10293393",
		DateOfAdmission: time.Now(),
	}
	student.details()
}

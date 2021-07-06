package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	entryNewStudent := true
	for entryNewStudent {
		var name, lastName, dni, exit string

		student := Student{}
		fmt.Println("Please, entry your name")
		fmt.Scanln(&name)

		fmt.Println("Please, entry your lastName")
		fmt.Scanln(&lastName)

		fmt.Println("Please, entry your dni")
		fmt.Scanln(&dni)

		student.setName(name)
		student.setLastName(lastName)
		student.setDni(dni)
		student.setAdmission(time.Now())

		student.details()

		fmt.Println("Do you want to continue attracting students? Put y/n")
		fmt.Scanln(&exit)

		if strings.ToLower(exit) == "n" {
			entryNewStudent = false
		}
	}
}

type Student struct {
	Name      string
	LastName  string
	Dni       string
	Admission time.Time
}

func (s *Student) setName(name string) {
	s.Name = name
}

func (s *Student) setLastName(lastName string) {
	s.LastName = lastName
}

func (s *Student) setDni(dni string) {
	s.Dni = dni
}

func (s *Student) setAdmission(admission time.Time) {
	s.Admission = admission
}

func (s Student) details() {
	fmt.Printf("Name: [%s]\n", s.Name)
	fmt.Printf("LastName: [%s]\n", s.LastName)
	fmt.Printf("Dni: [%s]\n", s.Dni)
	fmt.Printf("Admission: [%v]\n", s.Admission)
}

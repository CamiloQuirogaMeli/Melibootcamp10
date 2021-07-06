package main

import (
	f "fmt"
	r "reflect"
)

var (
	salary float32
	tax    float32
)

type Date struct {
	day   int
	month int
	year  int
}

type Student struct {
	Name           string
	Surname        string
	DNI            int
	Admission_date Date
}

func main() {

	menu := "Do you want to register another student? (1 yes: 2 no): "
	i := 1
	student := Student{}

	for i == 1 {
		f.Print("Enter the student's name: ")
		f.Scanln(&student.Name)

		f.Print("Enter the student's surname: ")
		f.Scanln(&student.Surname)

		f.Print("Enter the student's DNI: ")
		f.Scanln(&student.DNI)

		f.Println("Enter the student's admission date: ")
		f.Print("Day: ")
		f.Scanln(&student.Admission_date.day)
		f.Print("Month: ")
		f.Scanln(&student.Admission_date.month)
		f.Print("Year: ")
		f.Scanln(&student.Admission_date.year)

		student.getDetails()

		f.Print(menu)
		f.Scanln(&i)
	}
}

func (s Student) getDetails() {

	values := r.ValueOf(s)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		f.Println(types.Field(i).Name, ":", values.Field(i))
	}
}

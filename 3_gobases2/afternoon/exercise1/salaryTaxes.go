package main

import "fmt"

type Student struct {
	firstName      string
	lastName       string
	dni            string
	onboardingDate string
}

func (student Student) detail() string {
	return "First name: " + student.firstName + "\nLast name: " + student.lastName + "\nDNI: " + student.dni + "\nOnboarding date: " + student.onboardingDate
}

func main() {
	student1 := Student{"Andres", "Moya", "123444", "15/06/21"}
	fmt.Println("This is the information about our first student:\n" + student1.detail())
	student2 := Student{
		firstName:      "Felipe",
		lastName:       "Rodríguez",
		dni:            "68767",
		onboardingDate: "15/16/21"}
	fmt.Println("This is the information about our second student:\n" + student2.detail())
}

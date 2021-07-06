package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	dni       string
	entryDate string
}

func (a Student) detalle() {
	fmt.Println("First Name:", a.firstName)
	fmt.Println("Last Name:", a.lastName)
	fmt.Println("DNI:", a.dni)
	fmt.Println("Entry Date:", a.entryDate)
}

func main() {

	a := Student{"Rodrigo", "Peró", "39526888", "22/06/2021"}

	a.detalle()
}

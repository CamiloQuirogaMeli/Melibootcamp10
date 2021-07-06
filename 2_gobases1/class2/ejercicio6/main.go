package main

import "fmt"

var employees = map[string]int{
	"Benjamin": 20,
	"Nahuel":   26,
	"Brenda":   19,
	"Darío":    44,
	"Pedro":    30,
}

func main() {

	infoOfEmployee("Benjamin")
	employeesOver21()
	addEmployee("Federico", 25)
	removeEmployee("Pedro")
}

func infoOfEmployee(name string) {

	var age int

	for key, value := range employees {
		if key == name {
			age = value
			break
		}
	}

	fmt.Println(name, age)
}

func addEmployee(name string, age int) {

	employees[name] = age

	fmt.Println(employees)
}

func removeEmployee(name string) {

	delete(employees, name)

	fmt.Println(employees)
}

func employeesOver21() {

	over21 := 0

	for _, value := range employees {
		if value > 21 {
			over21++
		}
	}

	fmt.Println(over21)
}

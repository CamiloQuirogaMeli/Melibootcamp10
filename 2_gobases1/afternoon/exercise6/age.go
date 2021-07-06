package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	queryEmployee, newEmployee, removedEmployee := "Benjamin", "Federico", "Pedro"

	fmt.Println("How old is Benjamin?\nHe is", employees[queryEmployee], "years old")
	fmt.Println("How many employees are over 21 years old?")
	fmt.Println(employeesOver21(employees), "employees are over 21 years old")

	employees[newEmployee] = 25
	fmt.Println("new employee added, named", newEmployee, ", who has", employees[newEmployee])

	delete(employees, removedEmployee)
	fmt.Println(removedEmployee, "is no longer working with us")

}

func employeesOver21(employees map[string]int) int {
	over21 := 0
	for _, age := range employees {
		if age > 21 {
			over21++
		}
	}
	return over21
}

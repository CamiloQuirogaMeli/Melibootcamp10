package main

import (
	"fmt"
)

var (
	employees = map[string]int{
		"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Dario": 44, "Pedro": 30}

	employeesOverTwentyOne []string
	newEmployeeAge         = 25
	newEmployeeName        = "Federico"
	employeeToDelete       = "Pedro"
)

func main() {

	fmt.Println("Part 0 - ", employees)

	employeesOverTwentyOne = getOverTwenteeOne(employees)

	fmt.Println("Part 1 - ", employeesOverTwentyOne)

	employees = addNewEmployee(employees, newEmployeeName, newEmployeeAge)

	fmt.Println("Part 2 - ", employees)

	employees = deleteEmployee(employees, employeeToDelete)

	fmt.Println("Part 3 - ", employees)
}

func getOverTwenteeOne(employees map[string]int) []string {

	var employeesToReturn []string

	for name, age := range employees {
		if age > 21 {
			employeesToReturn = append(employeesToReturn, name)
		}
	}

	return employeesToReturn
}

func addNewEmployee(employees map[string]int, name string, age int) map[string]int {
	employees[name] = age
	return employees
}

func deleteEmployee(employees map[string]int, employee string) map[string]int {
	delete(employees, employee)
	return employees
}

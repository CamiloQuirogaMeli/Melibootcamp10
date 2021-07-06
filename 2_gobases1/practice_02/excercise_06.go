package main

import "fmt"

const (
	LEGAL_AGE = 21
)

func main() {

	var employeeName string = "Benjamin"
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Age %s : [%d]\n", employeeName, employees[employeeName])

	var employeesLegalAge uint = 0

	for _, age := range employees {
		if age >= LEGAL_AGE {
			employeesLegalAge++
		}
	}

	fmt.Printf("Employees of legal age: [%d]\n", employeesLegalAge)

	var newEmployee string = "Federico"
	fmt.Printf("Added employee: [%s]\n", newEmployee)
	employees[newEmployee] = 25

	var deleteEmployee string = "Pedro"
	fmt.Printf("Delete employee: [%s]\n", deleteEmployee)
	delete(employees, "Pedro")

	fmt.Println(employees)
}

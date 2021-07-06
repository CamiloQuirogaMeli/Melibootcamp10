package main

import (
	f "fmt"
)

var (
	salary float32
	tax    float32
)

func main() {

	f.Print("Enter Salary: $")
	f.Scanln(&salary)

	f.Println("Salary tax $", calculateTax(salary))

}

func calculateTax(salary float32) float32 {

	if salary > 150000 {
		tax = salary*17/100 + (salary-salary*17/100)*10/100
	} else if salary > 50000 {
		tax = salary * 17 / 100
	}

	return tax
}

package main

// You need run go get github.com/leekchan/accounting for correct money and currency format
import (
	f "fmt"
	"github.com/leekchan/accounting"
)

func taxSalary(salary float32) float32 {
	var tax float32
	var salaryTax float32
	switch {
	case salary > 150000:
		tax = 27
	case salary > 50000:
		tax += 10
	default:
		tax = 0
	}
	salaryTax = salary * (tax / 100)
	return salaryTax
}

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	var salary float32
	var salaryTax float32
	f.Println("Welcome, enter your salary to know your tax salary: ")
	f.Scanln(&salary)
	salaryTax = taxSalary(salary)
	f.Println("Your tax salary", ac.FormatMoney(salaryTax))
}

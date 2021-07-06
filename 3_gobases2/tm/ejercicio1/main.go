package main

import "fmt"

func getNetSalary(grossSalary float64) (netSal float64) {
	var tax float64
	if grossSalary > 50000 {
		tax = 17
	}
	if grossSalary > 150000 {
		tax += 10
	}
	if tax != 0 {
		toSub := (grossSalary * tax) / 100
		netSal = grossSalary - toSub
		return netSal
	}
	return grossSalary
}

func main() {
	var salary float64 = 60000
	netSalary := getNetSalary(salary)
	fmt.Printf("Salary to pay: $%.2f\n", netSalary)
}

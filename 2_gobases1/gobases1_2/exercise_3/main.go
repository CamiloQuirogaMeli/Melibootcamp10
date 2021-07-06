package main

import "fmt"

var (
	age              int
	employee         string
	months_seniority int
	salary           float32
	loan             bool
	interest_rate    float32
)

func main() {

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Is employed? (yes or no): ")
	fmt.Scanln(&employee)

	fmt.Print("Enter the months of seniority working: ")
	fmt.Scanln(&months_seniority)

	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)

	if age < 22 {
		fmt.Print("\nIt is not possible to grant a loan since you must be over 22 years old")
	} else if employee == "no" {
		fmt.Print("\nIt is not possible to grant a loan since you are unemployed")
	} else if months_seniority < 12 {
		fmt.Print("\nIt is not possible to grant a loan since you do not have more than one year seniority working")
	} else {
		if salary > 100000 {
			interest_rate = 0
		} else {
			interest_rate = 1.8
		}
		fmt.Printf("\nYou are granted a loan with an interest rate of %.1f %%", interest_rate)
	}
}

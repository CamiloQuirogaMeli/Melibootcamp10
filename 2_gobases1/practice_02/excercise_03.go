package main

import "fmt"

const (
	MINIMUM_AGE_REQUIRED       = 22
	LABOR_OLD_REQUIRED         = 1
	LOAN_WITHOUT_INTEREST_FROM = 100000.00
)

func main() {

	var age int
	fmt.Println("Please, input age")
	fmt.Scanln(&age)

	var isEmployee bool
	fmt.Println("Is employee? Please, indicate true or false")
	fmt.Scanln(&isEmployee)

	var laborOld int
	fmt.Println("Enter the work seniority")
	fmt.Scanln(&laborOld)

	var salary float64
	fmt.Println("Please, enter the salary")
	fmt.Scanln(&salary)

	if age > MINIMUM_AGE_REQUIRED {
		if isEmployee && laborOld > LABOR_OLD_REQUIRED {
			if salary > LOAN_WITHOUT_INTEREST_FROM {
				fmt.Println("Congrulations, you have an interest-free loan!")
			} else {
				fmt.Println("Congrulations, you have a loan!.")
			}
		} else {
			fmt.Println("Sorry, you don't have some requirements")
		}
	} else {
		fmt.Println("Sorry, you do not have the minium age required.")
	}

}

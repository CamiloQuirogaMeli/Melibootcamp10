package main

import (
	"fmt"
)

func main() {
	var age int = 0
	var salary float32 = 0.0
	var seniority float32 = 0.0
	var is_employee bool = true
	var interest_rate bool = true

	fmt.Println("...Welcome to national bank...")
	fmt.Println("...If you want to know if you are elegible for a loan, enter the following info...")
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	fmt.Print("Working time (In years): ")
	fmt.Scanln(&seniority)
	fmt.Print("Are you an employee? 1 = yes, 0 = no: ")
	fmt.Scanln(&is_employee)
	fmt.Print("Salary: ")
	fmt.Scanln(&salary)

	if age > 22 && is_employee && seniority > 1 {
		if salary > 100000 {
			interest_rate = !interest_rate
		}
		if interest_rate {
			fmt.Println("Your elegible for a loan with interest rate!")
		}else{
			fmt.Println("Your elegible for a loan without interest rate !")
		}
	}else{
		fmt.Println("Sorry, you are not elegible for a loan")
	}
}

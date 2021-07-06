package main

import "fmt"

func main() {

	var age, seniorityAge int
	var employed string
	var salary float64

	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age < 22 {
		fmt.Println("You have to be at least 22 years old.")
		return
	}

	fmt.Println("Are you employed? (Y/N):")
	fmt.Scanln(&employed)

	if employed != "Y" && employed != "y" {
		fmt.Println("You have to be employed.")
		return
	}

	fmt.Println("Enter your seniority age:")
	fmt.Scanln(&seniorityAge)

	if seniorityAge < 1 {
		fmt.Println("You must be at least 1 year of seniority.")
		return
	}
	fmt.Println("Credit granted!")

	fmt.Println("Enter your salary:")
	fmt.Scanln(&salary)

	if salary > 100000 {
		fmt.Println("The credit will not have interest")
	} else {
		fmt.Println("The credit will have interest")
	}

}

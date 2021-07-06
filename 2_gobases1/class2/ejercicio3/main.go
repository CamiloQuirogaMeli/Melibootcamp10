package main

import "fmt"

func main() {
	var age int = 22
	var employed bool = true
	var salary float64 = 100000
	var monthsInCompany int = 13

	loan(age, employed, salary, monthsInCompany)
}

func loan(age int, employed bool, salary float64, monthsInCompany int) {
	switch {
	case age > 22 && employed == true && monthsInCompany > 12 && salary > 100000:
		fmt.Println("Congratulations! You're applicable to receive a loan without interests.")
	case age > 22 && employed == true && monthsInCompany > 12 && salary <= 100000:
		fmt.Println("Congratulations! You're applicable to receive a loan but you will have to pay interests.")
	default:
		fmt.Println("Sorry! You're not applicable to get a loan.")
	}
}

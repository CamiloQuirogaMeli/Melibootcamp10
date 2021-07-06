package main

import "fmt"

func isEligibleForDiscount(age int, employed bool, seniorityYears int) bool {
	return age > 22 && employed && seniorityYears > 1 
}

func main() {
	var age int = 23
	var employed bool = true
	var seniorityYears int = 5
	var income int = 100001

	if isEligibleForDiscount(age, employed, seniorityYears) {
		var text string
		if income > 100000 {
			text = "sin intereses"
		} else {
			text = "con intereses"
		}
		fmt.Println("El empleado puede acceder a un prestamo", text)
	} else {
		fmt.Println("El empleado no puede acceder a un prestamo")
	}
}
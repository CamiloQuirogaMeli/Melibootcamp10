package main

import "fmt"

const PAY_CATEGORY_C = 1000.0
const BONUS_CATEGORY_B = 20.0
const PAY_CATEGORY_B = 1500.0
const BONUS_CATEGORY_A = 50.0
const PAY_CATEGORY_A = 3000.0

// salaryByMonth return the salary assumming that work 8hs per day
func salaryByMonth(salaryByH, bonus float64) float64 {
	grossSalary := 160 * salaryByH

	return grossSalary - ((bonus * grossSalary) / 100)
}

// getSalary returns the salary by minutes of work and category of the employee
func getSalary(workByMinutes float64, category string) float64 {
	hoursWorked := workByMinutes / 60
	var salary float64

	switch category {
	case "C":
		salary = (PAY_CATEGORY_C / 1) * hoursWorked
	case "B":
		salary = (PAY_CATEGORY_B / hoursWorked) + salaryByMonth(PAY_CATEGORY_B, BONUS_CATEGORY_B)
	case "A":
		salary = (PAY_CATEGORY_A / hoursWorked) + salaryByMonth(PAY_CATEGORY_A, BONUS_CATEGORY_A)
	}
	return salary
}

func main() {
	var workByMinutes float64 = 480
	var category string = "B"
	salary := getSalary(workByMinutes, category)
	fmt.Printf("%.2f\n", salary)
}

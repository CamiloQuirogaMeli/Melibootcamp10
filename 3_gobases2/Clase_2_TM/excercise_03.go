package main

import (
	"fmt"
	"strings"
)

const (
	CATEGORY_C = 1000
	CATEGORY_B = 1500
	CATEGORY_A = 3000
)

func main() {

	var minutesWorked uint
	var category string

	fmt.Println("Please, enter your minutes worked")
	fmt.Scanln(&minutesWorked)

	fmt.Println("Please, enter your category letter")
	fmt.Scanln(&category)

	salary := getMonthlySalaryByEmployeeCategory(minutesWorked, strings.ToUpper(category))
	fmt.Printf("You salary is: [%2.f]\n", salary)
}

func getMonthlySalaryByEmployeeCategory(minutesMonthly uint, category string) float64 {
	workHours := minutesMonthly / 60

	switch category {
	case "C":
		return salaryCategoryC(workHours)
	case "B":
		return salaryCategoryDistinctC(workHours, CATEGORY_B, 0.20)
	case "A":
		return salaryCategoryDistinctC(workHours, CATEGORY_A, 0.50)
	}

	return 0
}

func salaryCategoryC(workHours uint) float64 {
	return float64(workHours * CATEGORY_C)
}

func salaryCategoryDistinctC(workHours uint, category uint, bondPercentage float64) float64 {
	monthlySalary := workHours * category
	return float64(monthlySalary) + (float64(monthlySalary) * bondPercentage)
}

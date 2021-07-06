package main

import "fmt"

func main() {
	workedHours := 95
	semesterSalaries := []int{79000, 92000, 79500, 85000, 0, 92000}
	salary, monthErr := monthlySalary(workedHours)
	bonus, bonusErr := bonus(semesterSalaries)

	if monthErr != nil {
		fmt.Println(monthErr)
	} else if bonusErr != nil {
		fmt.Println(bonusErr)
	} else {
		fmt.Println("Your salary is", salary)
		fmt.Println("Your bonus is", bonus)
	}
}

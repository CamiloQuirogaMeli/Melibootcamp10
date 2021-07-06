package main

import "fmt"

func main() {
	hourlyWage := 15000.
	bestSemesterSalary := 140000.

	monthlySalary_1, e1 := calculateSalary(81, hourlyWage)
	handleErrorSalary(monthlySalary_1, e1)

	monthlySalary_2, e2 := calculateSalary(-5, hourlyWage)
	handleErrorSalary(monthlySalary_2, e2)

	fmt.Println()

	semiannualBonus, e3 := calculateSemiannualBonus(bestSemesterSalary, 4)
	handleErrorBonus(semiannualBonus, e3)

	semiannualBonus, e4 := calculateSemiannualBonus(bestSemesterSalary, -3)
	handleErrorBonus(semiannualBonus, e4)

}

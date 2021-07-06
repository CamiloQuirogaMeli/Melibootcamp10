package main

import "fmt"

func employeeIncome(monthWorkHours int, category string) float64 {
	var income float64 = 0
	switch category {
		case "C":
			income = 1000 * float64(monthWorkHours)
		case "B":
			income = 1500 * float64(monthWorkHours)
			income += 0.2 * income
		case "A":
			income = 3000 * float64(monthWorkHours)
			income += 0.5 * income
	}
	return income
}

func main() {

	
	fmt.Println("* Caso 1: categoria A")
	fmt.Printf("\tHoras por mes: %d\n\tSalario: %.2f\n", 30, employeeIncome(30, "A"))

	fmt.Println("* Caso 2: categoria B")
	fmt.Printf("\tHoras por mes: %d\n\tSalario: %.2f\n", 40, employeeIncome(40, "B"))

	fmt.Println("* Caso 3: categoria C")
	fmt.Printf("\tHoras por mes: %d\n\tSalario: %.2f\n", 50, employeeIncome(50, "C"))
}

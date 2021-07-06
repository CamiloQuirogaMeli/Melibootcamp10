package main

import "fmt"

func calculateSalary(minutesWorked int, category string) float64 {
	var salary float64
	switch category {
	case "A":
		salary = float64(int(minutesWorked/60)*3000) * 1.5
	case "B":
		salary = float64(int(minutesWorked/60)*1500) * 1.2
	case "C":
		salary = float64(int(minutesWorked/60) * 1000)
	}

	return salary
}

func main() {
	var salary = calculateSalary(120, "A")
	fmt.Printf("Salario es %.2f\n", salary)
}

package main

import (
	"fmt"
)

func main() {
	var minutes float64
	var category string

	fmt.Println("Enter the minutes:")
	fmt.Scanln(&minutes)

	fmt.Println("Enter the category:")
	fmt.Scanln(&category)

	fmt.Println("Salary:", calcSalary(minutes, category))

}

func calcSalary(minutes float64, category string) float64 {

	hours := float64(minutes) / 60.0

	switch category {
	case "C":
		return hours * 1000
	case "B":
		return (hours * 1500) * 1.2
	case "A":
		return (hours * 3000) * 1.5
	default:
		return 0
	}

}

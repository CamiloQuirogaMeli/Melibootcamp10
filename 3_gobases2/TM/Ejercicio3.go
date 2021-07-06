package main

import "fmt"

const (
	CATEGORIE_A = "A"
	CATEGORIE_B = "B"
	CATEGORIE_C = "C"
)

func main() {
	result := processedSalary(5000, "A")
	fmt.Println(result)
}

func processedSalary(minutesPerMonth int, category string) float64 {
	hours := minutesToHours(minutesPerMonth)
	var salary float64
	switch category {
	case CATEGORIE_A:
		salary = (3000 * hours) * 1.5
	case CATEGORIE_B:
		salary = (1500 * hours) * 1.2
	case CATEGORIE_C:
		salary = 1000 * hours
	}
	return salary
}

func minutesToHours(value int) float64 {
	return float64(value) / 60
}

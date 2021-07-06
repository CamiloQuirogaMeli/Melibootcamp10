package main

import (
	"fmt"
	"math"
	"strings"
)

func calculateSalary(minutesWorked int, category string) float64 {

	var hoursWorked int = int(math.Floor(float64(minutesWorked / 60)))

	category = strings.ToUpper(category)

	switch category {
	case "C":
		return float64(hoursWorked * 1000)
	case "B":
		return float64((hoursWorked * 1500) + ((hoursWorked * 1500) * 20 / 100))
	case "A":
		return float64((hoursWorked * 3000) + ((hoursWorked * 1500) * 50 / 100))
	}

	return float64(hoursWorked * 1000)
}

func main() {

	workerSalary := calculateSalary(4000, "A")

	fmt.Println(workerSalary)
}

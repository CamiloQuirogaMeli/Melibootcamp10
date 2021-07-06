package main

import "fmt"

func calculateSalary(minutes_per_month int, category string) float64 {
	hoursPerMonth := float64(minutes_per_month) / 60.0
	hourlyPriceByCategory := map[string]float64{"A": 3000, "B": 1500, "C": 1000}
	bonusByCategory := map[string]float64{"A": 0.5, "B": 0.2, "C": 0}
	return hoursPerMonth * hourlyPriceByCategory[category] * (1 + bonusByCategory[category])
}

func main() {
	minutesPerMonth := 6000
	category := "A"
	fmt.Println("El salario del empleado que trabajó", minutesPerMonth, "minutos",
		"y es categoría", category, "es: "+"$"+fmt.Sprint((calculateSalary(minutesPerMonth, category))))
}

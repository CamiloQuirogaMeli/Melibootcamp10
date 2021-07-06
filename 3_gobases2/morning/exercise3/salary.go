package main

import "fmt"

const (
	A string = "A"
	B string = "B"
	C string = "C"
)

func main() {
	workedMinutes := 10000
	category := "C"
	valuePerHour := map[string]int{A: 3000, B: 1500, C: 1000}
	bonusPercentage := map[string]float32{A: 0.5, B: 0.2, C: 0.0}
	salary := getSalary(workedMinutes, category, valuePerHour, bonusPercentage)
	fmt.Println("Given that you worked", workedMinutes, "this month and you have", category, "category, your salary is", salary)

}

func getSalary(workedMinutes int, category string, valuePerHour map[string]int, bonusPercentage map[string]float32) float32 {

	var workedHours float32 = float32(workedMinutes) / 60.0

	return workedHours * float32(valuePerHour[category]) * (1 + bonusPercentage[category])
}

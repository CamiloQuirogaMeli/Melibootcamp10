package main

import (
	f "fmt"
)

var (
	minutes  int32
	category string
	salary   float32
)

func main() {

	f.Print("Enter the minutes worked per month: ")
	f.Scanln(&minutes)

	f.Print("Enter category: ")
	f.Scanln(&category)

	f.Println(calculateSalary(minutes, category))
}

func calculateSalary(minutes int32, category string) float32 {

	switch category {
	case "A":
		salary = 1000 * float32(minutes) / 60
	case "B":
		salary = 1500 * float32(minutes) / 60
		salary = salary + 0.2*salary
	case "C":
		salary = 3000 * float32(minutes) / 60
		salary = salary + 0.5*salary
	}
	return salary
}

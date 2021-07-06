package main

import (
	"fmt"
	"strings"
)

func main() {

	var category string
	var time float64
	var salary float64

	fmt.Println("Ingrese la categoria del empleado: \n")
	fmt.Scanf("%s", &category)

	fmt.Println("Ingrese la categoria de horas trabajadas: \n")
	fmt.Scanf("%f", &time)

	salary = getSalary(time, category)

	fmt.Printf("El salario de este mes es de $%.2f ", salary)

}

func getSalary(time float64, category string) float64 {

	var salary float64
	category = strings.ToUpper(category)
	time = time * 60

	switch category {
	case "A":
		salary = (getSalaryForMinute(category) * time) * 1.5
	case "B":
		salary = (getSalaryForMinute(category) * time) * 1.2
	case "C":
		salary = (getSalaryForMinute(category) * time)

	}
	return salary

}

func getSalaryForMinute(category string) float64 {
	var minuteVal float64

	switch category {

	case "A":
		minuteVal = 3000.0 / 60.0

	case "B":
		minuteVal = 1500.0 / 60.0

	case "C":
		minuteVal = 1000.0 / 60.0

	}

	return minuteVal

}

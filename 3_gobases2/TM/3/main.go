package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	var minutes int
	var category string

	fmt.Println("Ingrese cantidad de minutos trabajados:")
	fmt.Scanln(&minutes)

	fmt.Println("Ingrese categoria de empleado:")
	fmt.Scanln(&category)

	monthlySalary, err := calculateSalary(minutes, category)

	if err == nil {
		fmt.Println("El salario mensual del empleado es: ", monthlySalary)
	} else {
		fmt.Println(err)
	}
}

func calculateSalary(minutes int, category string) (float64, error) {
	var salary int
	var additional float64

	switch strings.ToUpper(category) {
	case "C":
		salary = 1000
	case "B":
		salary = 1500
		additional = 20
	case "A":
		salary = 3000
		additional = 50
	default:
		return 0, errors.New("Categoria erronea")
	}
	monthlySalary := float64((minutes / 60) * salary)

	if additional > 0 {
		monthlySalary += float64((additional * monthlySalary) / 100)
	}

	return monthlySalary, nil
}

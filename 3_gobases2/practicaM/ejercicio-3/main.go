package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := calculateSalary(1000, "A")

	if err == nil {
		fmt.Println("El salario del empleado es:", salary)
	} else {
		fmt.Println(err)
	}
}

func calculateSalary(minPerMonth int, category string) (float64, error) {

	if minPerMonth <= 0 {
		return 0, errors.New("ERROR! Los minutos trabajados no pueden ser negativos, tampoco cero")
	}

	var hoursPerMonth float64 = float64(minPerMonth) / 60

	switch category {
	case "C":
		return hoursPerMonth * 1000, nil
	case "B":
		salary := hoursPerMonth * 1500
		return salary + (salary * 0.2), nil
	case "A":
		salary := hoursPerMonth * 3000
		return salary + (salary * 0.5), nil
	default:
		return 0, errors.New("ERROR! Categoria inexistente")
	}
}

package main

import (
	"errors"
	"fmt"
)

func calculateSalary(workingTime int) (float64, error) {
	if workingTime < 80 {
		return 0, errors.New("EEROR! El trabajador no puede haber trabajado menos de 80 hs mensuales.")
	}
	salary := float64(workingTime * 200)
	if salary < 150000 {
		return salary, nil
	}
	return salary * 0.9, nil
}

func calculateBonus(betterSalary float64, workingMonths int) (float64, error) {
	if betterSalary < 0 {
		return 0, errors.New("ERROR: El salario no puede ser negativo")
	}
	return (betterSalary / 12) * float64(workingMonths), nil
}

func main() {
	var hours int = 90
	sal, err := calculateSalary(hours)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("salario: ", sal)

	var salary float64 = 1013
	var months int = 5
	bonus, err := calculateBonus(salary, months)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("Aguinaldo: %.2f\n", bonus)
}

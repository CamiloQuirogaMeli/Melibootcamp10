package main

import (
	"errors"
	"fmt"
)

func calculateSalary(salary, workedHours int) (float64, error) {

	if workedHours < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
	}

	var totalSalary float64

	if salary >= 150.000 {
		totalSalary = float64(salary + (salary * 10 / 100))
		return totalSalary, nil
	}

	if salary < 0 {
		return float64(salary), nil
	} else {
		e := fmt.Sprintf("error: el salario no puede ser menos de 0. salario ingresado: %d", salary)
		return 0, errors.New(e)
	}
}

func calculateBonus(salaries []float64, monthsWorked int) (float64, error) {

	if monthsWorked <= 0 {
		return 0, fmt.Errorf("error: los meses trabajados no pueden ser menos de 0")
	}

	var bestSalary float64
	for _, salary := range salaries {

		if bestSalary < salary {
			bestSalary = salary
		}
	}

	bonus := (bestSalary / 12) * float64(monthsWorked)
	return bonus, nil
}

func main() {

	salary, hoursWorked := 160000, 80

	totalSalary, err1 := calculateSalary(salary, hoursWorked)

	fmt.Println(totalSalary)

	salaries := []float64{100, 200, 300, 400, 500, 600}
	monthsWorked := len(salaries)

	totalBonus, err2 := calculateBonus(salaries, monthsWorked)

	errorss := fmt.Errorf("%w and %w", err1, err2)

	var err error
	err = errors.Unwrap(errorss)
	fmt.Println(err)
	err3 := errors.Unwrap(err)
	fmt.Println(err3)

	fmt.Println(totalBonus)
}

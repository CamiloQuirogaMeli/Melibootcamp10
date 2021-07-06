package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	msg string
}

func (e SalaryError) Error() string {
	return "Error de salario ocurrido. Favor presentar a RRHH"
}

func calculateSalary(salary float64, hoursWorked int) (float64, error) {

	if hoursWorked < 80 {
		return 0, SalaryError{
			msg: "el trabajador no puede haber trabajado menos de 80 hs mensuales",
		}
	}

	if salary >= 150000 {
		return salary * (1 - 0.1), nil
	} else {
		return salary, nil
	}
}

func main() {
	finalSalary, err := calculateSalary(150000, 70)
	e1 := fmt.Errorf("SalaryError: %w", err)
	if errors.Unwrap(e1) != nil {
		fmt.Println(errors.Unwrap(e1))
		fmt.Println(errors.New("Revise nuevamente su salario por si tiene que cambiar"))
	} else {
		fmt.Printf("Salario final es $%.2f", finalSalary)
	}
}

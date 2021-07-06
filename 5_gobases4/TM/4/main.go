package main

import (
	"errors"
	"fmt"
)

func main() {
	var hours int
	fmt.Println("Ingrese las horas trabajadas:")
	fmt.Scanln(&hours)

	var salary, err = GetSalary(hours)

	if err == nil {
		fmt.Println("El salario es ", salary)

		var months int
		fmt.Println("Ingrese los meses trabajados:")
		fmt.Scanln(&months)

		amount, err2 := GetAmount(salary, months)

		if err2 == nil {
			fmt.Println("El aguinaldo es ", amount)
		} else {
			fmt.Print(errors.Unwrap(err2))
		}
	} else {
		fmt.Print(err)
	}
}

func GetSalary(hours int) (float64, error) {
	if hours < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salary := float64(hours * 1000)

	if salary >= 150000 {
		salary -= (salary * 0.10)
	}

	return salary, nil
}

func GetAmount(salary float64, months int) (float64, error) {
	if months < 0 || salary < 0 {

		return 0, fmt.Errorf("error interno: %w", internError{})
	}

	var amount = (salary / 12) * float64(months)
	return amount, nil
}

type internError struct{}

func (i internError) Error() string {
	return "error: el valor ingresado es negativo"
}

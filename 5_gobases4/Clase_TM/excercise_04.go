package main

import (
	// "errors"
	"fmt"
	"os"
)

func main() {
	var salary, betterSalary float64
	var workHours, monthsWorked int

	fmt.Println("Please, entry your salary")
	fmt.Scanln(&salary)

	fmt.Println("Please, entry your work hours")
	fmt.Scanln(&workHours)

	pay, err := monthlySalary(salary, workHours)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Your pay is: [%2.f]\n", pay)

	fmt.Println("Please, entry your betterSalary")
	fmt.Scanln(&betterSalary)

	fmt.Println("Please, entry your months worked")
	fmt.Scanln(&monthsWorked)

	bonus, err := calculateBonus(betterSalary, monthsWorked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Your bonus is: [%2.f]\n", bonus)
}

func monthlySalary(salary float64, workHours int) (float64, error) {
	if workHours < 80 {
		return 0, &SalaryError{
			msg: "el trabajador no puede haber trabajado menos de 80hs mensuales",
		}
	}
	//Deberia chequear cuanto es el salario mensual por las 80 hs y si se pasaan
	if salary >= 150000 {
		discount := salary * 0.10
		return salary - discount, nil
	}

	return salary, nil
}

func calculateBonus(betterSalary float64, monthsWorked int) (float64, error) {

	if monthsWorked == 0 {
		return 0.0, &SalaryError{
			msg: "la cantidad de meses trabajados no puede ser cero",
		}
	}

	if monthsWorked > 6 {
		return 0.0, &SalaryError{
			msg: "la cantidad de meses trabajados no puede superar un semestre (6 meses)",
		}
	}

	if betterSalary < 0.0 || monthsWorked < 0 {
		return 0.0, &SalaryError{
			msg: "el salario o la cantidad de meses trabajados no pueden ser negativos",
		}
	}

	bonus := (betterSalary / 12) * float64(monthsWorked)
	return bonus, nil
}

type SalaryError struct {
	msg string
}

func (e SalaryError) Error() string {
	return fmt.Sprintf("error: %s", e.msg)
}

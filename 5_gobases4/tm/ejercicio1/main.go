package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type SalaryError struct {
	Salary int
	Err    error
}

func (se *SalaryError) Error() string {
	return fmt.Sprintf("the salary entered, %d does not reach the taxable minimum", se.Salary)
}

func main() {
	var salary int = 145000
	if salary < 150000 {
		salaryErr := SalaryError{
			Salary: salary,
			Err:    nil,
		}
		log.Println(salaryErr)
		os.Exit(0)
	}
	fmt.Println("Must pay taxes")
}

// Para ejercicio 2
func ErrorsWithNew() error {
	return errors.New("the salary entered does not reach the taxable minimum")
}

// Para ejercicio 3
func ErrorsWithFmt(salary int) error {
	return fmt.Errorf("the taxable minimum is $150.000. The salary entered is %d", salary)
}

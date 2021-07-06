package main

import (
	"fmt"
	"os"
)

type salaryBandError struct {
	salary int
	msg    string
}

func salaryBandErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return salary, &salaryBandError{
			salary: salary,
			msg:    "debe pagar impuesto",
		}
	}
	return salary, nil
}

func (e *salaryBandError) Error() string {
	return fmt.Sprintf("el salario de $%d pesos %v", e.salary, e.msg)
}

func main() {
	var salary int = 140000

	salary, err := salaryBandErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

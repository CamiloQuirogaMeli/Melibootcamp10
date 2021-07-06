package main

import (
	"fmt"
	"os"
)

type minSalaryError struct {
	msg string
}

func (minError *minSalaryError) Error() string {
	return fmt.Sprintf("%s", minError.msg)
}

func minimumError(salary int) (string, error) {
	if salary < 150000 {
		return "", &minSalaryError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}

	return "Debe pagar impuesto", nil
}

func main() {
	var salary int = 1200000
	tax, err := minimumError(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(tax, "\n")
}

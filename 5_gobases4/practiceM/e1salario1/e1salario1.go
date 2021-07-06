package main

import (
	"fmt"
	"os"
)

type SalaryError struct {
	status int
	msg    string
}

func main() {

	salary := 42
	err := testSalary(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuestos")
}

func testSalary(s int) error {
	if s < 15000 {
		return &SalaryError{
			status: s,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf(("%d - %v"), e.status, e.msg)
}

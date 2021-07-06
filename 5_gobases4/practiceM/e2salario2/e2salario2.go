package main

import (
	"errors"
	"fmt"
	"os"
)

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
		return errors.New(fmt.Sprintf("%d - error: el salario ingresado no alcanza el mínimo imponible", s))
	}
	return nil
}

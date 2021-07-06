package main

import (
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
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", s)
	}
	return nil
}

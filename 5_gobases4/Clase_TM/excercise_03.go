package main

import (
	"fmt"
	"os"
)

func main() {

	var salary int

	fmt.Println("Please, entry your salary")
	fmt.Scanln(&salary)

	err := verifySalary(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

func verifySalary(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de: [%d]", salary)
	}
	return nil
}

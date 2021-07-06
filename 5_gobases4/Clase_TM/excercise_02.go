package main

import (
	"errors"
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
		return errors.New("error: el salario ingresado no alcanza el minimo imponible")
	}
	return nil
}

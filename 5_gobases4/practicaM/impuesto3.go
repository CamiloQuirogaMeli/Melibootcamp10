package main

import (
	"fmt"
	"os"
)

func myNewError(salary int) (string, error) {
	if salary < 150000 {
		err := fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		return "", err
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary, err := myNewError(15000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", salary)
}

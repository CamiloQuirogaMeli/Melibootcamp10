package main

import (
	"errors"
	"fmt"
	"os"
)

func minimumError(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}

	return "Debe pagar impuesto", nil
}

func main() {
	var salary int = 1222000
	tax, err := minimumError(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(tax)
}

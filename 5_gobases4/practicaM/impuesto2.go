package main

import (
	"errors"
	"fmt"
	"os"
)

func myNewError(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
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

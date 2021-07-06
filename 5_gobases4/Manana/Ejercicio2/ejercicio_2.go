package main

import (
	"errors"
	"fmt"
)

func checkSalary(salary int) bool {
	return salary < 150000
}

func main() {
	salary := 15000
	if checkSalary(salary) {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}

	fmt.Println("Debe pagar impuestos")
}

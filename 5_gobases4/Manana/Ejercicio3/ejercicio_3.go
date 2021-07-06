package main

import (
	"fmt"
)

func checkSalary(salary int) bool {
	return salary < 150000
}

func main() {
	salary := 15000
	if checkSalary(salary) {
		err := fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Debe pagar impuestos")
}

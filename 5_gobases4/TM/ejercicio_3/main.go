package main

import (
	"fmt"
	"os"
)

func IsSalaryInvalid(salary int) bool {
	
	return salary < 150000
}

func main(){

	var salary int = 100000

	if IsSalaryInvalid(salary) {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
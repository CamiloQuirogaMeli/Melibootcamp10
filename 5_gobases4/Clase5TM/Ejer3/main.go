package main

import (
	"fmt"
	"os"
)

func myCustomErrorTest(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}

func main() {
	var salary int = 149999
	err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto.\n")
}

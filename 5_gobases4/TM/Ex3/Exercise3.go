package main

import (
	"fmt"
	"os"
)

func errorFactory(salary int) error {

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		return err
	}
	return nil
}

func main() {

	salary := 160000

	err := errorFactory(salary)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

}

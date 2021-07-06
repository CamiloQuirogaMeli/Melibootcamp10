package ej2

import (
	"errors"
	"fmt"
)

func testSalario(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func ValidarSalario(salary int) {
	err := testSalario(salary)

	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	} else {
		fmt.Printf("Debe pagar impuesto")
	}
}

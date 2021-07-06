package main

import (
	"fmt"
)

type err interface {
	Error() error
}

type myError struct {
	message string
}

func (e myError) Error(salary int) error {
	if salary < 150000 {
		e.message = "ERROR! El mínimo imponible es de 150.000 y el salario ingresado es de: %d"
		return fmt.Errorf(e.message, salary)
	} else {
		e.message = "Debe pagar impuesto"
	}
	return fmt.Errorf(e.message)
}

func main() {
	var salary int = 1234
	e := myError{}
	fmt.Println(e.Error(salary))
}

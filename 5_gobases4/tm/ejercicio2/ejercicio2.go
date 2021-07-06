package main

import (
	"errors"
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
		e.message = "ERROR! El salario ingresado no alcanza el mínimo imponible"
	} else {
		e.message = "Debe pagar impuesto"
	}
	return errors.New(e.message)
}
func main() {
	var salary int = 1234
	e := myError{}
	fmt.Println(e.Error(salary))
}

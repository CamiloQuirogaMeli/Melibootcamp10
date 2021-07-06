package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type myCustomError struct {
	msg string
}

func (e myCustomError) Error(salary int) string {
	return e.msg
}

func main() {
	var salary int = 34500
	e := myCustomError{}

	e.Error(salary)

	if salary < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println("error ocurrido: ", err)
		return
	}
	e.msg = "debe pagar impuesto"
	fmt.Println("error ocurrido: ", e.msg)
}

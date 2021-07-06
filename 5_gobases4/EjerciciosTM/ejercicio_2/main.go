package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error(salary int) string {
	return e.msg
}

func main() {
	var salary int = 34500
	e := myCustomError{}

	e.Error(salary)

	if salary < 150000 {
		e.msg = "error: el salario ingresado no alcanza el minimo imponible"
	} else {
		e.msg = "debe pagar impuesto"
	}

	fmt.Println(errors.New(e.msg))
}

package main

import "fmt"

type error interface {
	Error() string
}

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error(salary int) string {
	if salary < 150000 {
		e.msg = "error: el salario ingresado no alcanza el minimo imponible"
	} else {
		e.msg = "debe pagar impuesto"
	}

	return e.msg
}

func main() {
	var salary int = 34500
	e := myCustomError{}
	fmt.Println(e.Error(salary))
}

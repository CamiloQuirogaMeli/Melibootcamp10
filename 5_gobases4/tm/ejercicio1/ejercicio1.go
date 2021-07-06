package main

import "fmt"

type error interface {
	Error() string
}

type myError struct {
	message string
}

func (err myError) Error(salary int) string {
	if salary < 150000 {
		err.message = "ERROR! El salario ingresado no alcanza el mínimo imponible"
	} else {
		err.message = "Debe pagar impuesto"
	}
	return err.message
}

func main() {
	var salary int = 1234
	err := myError{}
	fmt.Println(err.Error(salary))
}

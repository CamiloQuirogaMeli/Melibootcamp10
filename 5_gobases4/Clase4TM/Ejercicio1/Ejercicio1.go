package main

import "fmt"

type customError struct {
}

type CustomError interface {
	Error()
}

func (e *customError) Error() string {
	return fmt.Sprintf("El salario ingresado no alcanza el minimo")
}

func main() {

	salary := 160000

	if salary < 150000 {
		e := customError{}
		fmt.Print(e.Error())
	} else {
		fmt.Print("Debe pagar impuestos")
	}
}

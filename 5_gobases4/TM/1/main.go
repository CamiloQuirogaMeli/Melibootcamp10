package main

import (
	"fmt"
)

func main() {
	var salary int
	fmt.Println("Ingrese el sueldo:")
	fmt.Scanln(&salary)

	if salary < 150000 {
		err := salaryError{}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

type salaryError struct {
	msg string
}

func (e *salaryError) Error() string {
	return "error: el salario ingresado no alcanza el minimo imponible"
}

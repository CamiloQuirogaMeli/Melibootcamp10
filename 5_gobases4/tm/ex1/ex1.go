package main

import (
	"fmt"
	"os"
)

type Error interface {
	Error() string
}

type perError struct {
	salary int
	msg    string
}

func (e *perError) Error() string {
	return fmt.Sprintf("%d - %v", e.salary, e.msg)
}

func myCustomError(salary int) error {
	if salary < 150000 {
		return &perError{
			salary: salary,
			msg:    "error: el salario ingresado no alcanza el minimo imponible",
		}
	}
	return nil
}

func main() {
	var salary int = 145000
	err := myCustomError(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar imouesto")
}

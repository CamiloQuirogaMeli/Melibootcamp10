package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return e.msg
}

func myErrorTest(salary int) error {
	if salary < 150000 {
		return &myError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func main() {
	var salary int = 15700
	err := myErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

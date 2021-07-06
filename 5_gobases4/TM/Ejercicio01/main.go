package main

import (
	"fmt"
	"os"
)

type myError struct {
	msg string
}

type Error interface {
	Error() string
}

func (e *myError) Error() string {
	return e.msg
}

func impuesto(salary int) (string, error) {
	if salary < 150000 {
		return "", &myError{"Error: El salario ingresado no alcanza el minimo imponible"}
	} else {
		return "Debe Pagar impuestos", nil
	}
}

func main() {
	salary := 1500
	response, err := impuesto(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response)
}

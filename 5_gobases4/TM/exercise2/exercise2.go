package main

import (
	"errors"
	"fmt"
	"os"
)

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %v", e.msg)
}

func main() {
	salary := 10000
	err := errors.New("el salario ingresado no alcanza el mínimo imponible")
	if salary < 150000 {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto.")
}

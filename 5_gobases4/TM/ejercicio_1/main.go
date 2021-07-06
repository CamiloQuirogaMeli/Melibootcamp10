package main

import (
	"fmt"
	"os"
)

type InvalidValueError struct {
	msg		string
}

func (err *InvalidValueError) Error() string {

	return fmt.Sprintf("error: %v", err.msg)
}

func IsSalaryInvalid(salary int) bool {
	
	return salary < 150000
}

func main(){

	var salary int = 100000

	if IsSalaryInvalid(salary) {
		err := &InvalidValueError{"el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
package main

import (
	"fmt"
	"log"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf(e.msg)
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return &myError{msg: "error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}

func main() {
	salary := 15000
	err := checkSalary(salary)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Debe pagar impuestos")
}

package main

import (
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
	err := fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	if salary < 150000 {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto.")
}

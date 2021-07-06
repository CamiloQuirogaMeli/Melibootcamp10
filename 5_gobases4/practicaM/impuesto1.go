package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	message string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func myCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", &myCustomError{
			message: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}

	return fmt.Sprintf("Debe pagar impuesto ya que su salario es %d", salary), nil
}

func main() {
	salary, err := myCustomErrorTest(150001)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", salary)
}

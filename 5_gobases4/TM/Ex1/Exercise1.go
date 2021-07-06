package main

import (
	"fmt"
	"os"
)

type customError struct {
	msg string
}

func (e customError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func errorFactory() error {

	err := customError{"error: el salario ingresado no alcanza el mínimo imponible"}

	return err
}

func main() {

	salary := 120000

	if salary < 150000 {

		err := errorFactory()

		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

}

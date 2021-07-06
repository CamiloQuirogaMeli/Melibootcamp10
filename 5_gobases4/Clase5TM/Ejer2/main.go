package main

import (
	"errors"
	"fmt"
	"os"
)

func myCustomErrorTest(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible\n")
	}
	return nil
}

func main() {
	var salary int = 149999
	err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto.\n")
}

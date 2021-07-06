package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var salary int = 15700
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 155000
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

package main

import (
	"errors"
	"fmt"
)

func impuestoSalario(salary int, minimo int) (string, error) {
	if salary < minimo {
		return "0", errors.New("error: el salario ingresado no alcanza el minimo imponible")
	}
	return "Debe pagar impuestos", nil
}

func main() {
	result, err := impuestoSalario(100000, 150000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

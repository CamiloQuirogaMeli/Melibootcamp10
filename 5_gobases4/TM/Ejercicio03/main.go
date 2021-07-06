package main

import (
	"fmt"
	"os"
)

func impuesto(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("el minimo imponible es de 150.000 y el salario ingresado es de %d", salary)
	} else {
		return "Debe Pagar impuestos", nil
	}
}

func main() {
	salary := 1500
	response, err := impuesto(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response)
}

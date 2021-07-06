package main

import (
	"fmt"
	"os"
)

func main() {
	salary := 10000
	respuesta, err := pagarImpuesto(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(respuesta)
}
func pagarImpuesto(valor int) (string, error) {
	if valor < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", valor)
	} else {
		return "Debe pagar impuesto", nil
	}
}

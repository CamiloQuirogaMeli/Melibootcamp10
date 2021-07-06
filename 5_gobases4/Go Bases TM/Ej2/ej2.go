package main

import (
	"errors"
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
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuesto", nil
	}
}

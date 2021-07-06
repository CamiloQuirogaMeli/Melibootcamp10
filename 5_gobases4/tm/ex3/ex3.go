package main

import (
	"fmt"
)

func main() {
	var salary int = 145000
	if salary < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150,000 y el salario ingresado es de: %v", salary)
		fmt.Println("error ocurrido: ", err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

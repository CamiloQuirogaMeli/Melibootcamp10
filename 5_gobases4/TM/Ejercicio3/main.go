package main

import (
	"fmt"
)

func main() {
	salary := 125000

	if salary < 150000 {
		err := fmt.Errorf("error: el minimo no imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println("error ocurrido: ", err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

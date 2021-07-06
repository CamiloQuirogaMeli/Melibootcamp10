package main

import (
	"fmt"
)

const MIN_SALARY int = 150000

func main() {

	var salary int = 200

	if salary < MIN_SALARY {
		err := fmt.Errorf("error: el mínimo imponible es de %d y el salario ingresado es de: %d", MIN_SALARY, salary)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

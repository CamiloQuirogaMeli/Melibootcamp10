package main

import (
	"errors"
	"fmt"
)

const MIN_SALARY int = 150000

func main() {

	var salary int = 200

	if salary < MIN_SALARY {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

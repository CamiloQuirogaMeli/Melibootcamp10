package main

import (
	"fmt"
	"os"
)

func main() {
	var salary int = 1200

	if salary < 150000 {
		error := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(error)
		os.Exit(1)
	}

	fmt.Printf("Debe pagar impuesto\n")
}

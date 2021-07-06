package main

import (
	"fmt"
	"os"
)

func main() {
	minutes := 1789
	category := "A"

	salary, err := calcularSalario(minutes, category)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("El salario es:", salary)

}

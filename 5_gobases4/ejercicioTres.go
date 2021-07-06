package main

import (
	"fmt"
)

func main() {
	var salary int = 140000

	if salary < 150000 {
		err := fmt.Errorf("el salario de %d debe pagar impuestos", salary)
		fmt.Println(err)
	}
}

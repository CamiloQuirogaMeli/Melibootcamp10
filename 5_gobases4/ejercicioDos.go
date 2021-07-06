package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 140000

	if salary < 150000 {
		fmt.Println(errors.New("el salario debe pagar impuesto"))
		return
	}
}

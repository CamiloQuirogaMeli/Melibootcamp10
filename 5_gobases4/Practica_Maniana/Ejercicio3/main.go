package main

import (
	"fmt"
)

func impuestoSalario(salary int, minimo int) string {
	if salary < minimo {
		err := fmt.Errorf("error: el minimo imponible es de %d y el salario ingresado es de %d", minimo, salary)
		return fmt.Sprint(err)
	}
	return "Debe pagar impuestos"
}

func main() {
	result := impuestoSalario(100000, 150000)
	fmt.Println(result)

}

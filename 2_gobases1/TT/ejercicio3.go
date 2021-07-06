package main

import "fmt"

func main() {
	var age int8 = 22
	var isEmployee bool = true
	var yearsAntiguaty int8 = 2
	var salary float32 = 100000

	switch {
	case age > 22 && isEmployee && yearsAntiguaty > 1:
		if salary > 100000 {
			fmt.Println("Préstamo con interés")
		} else {
			fmt.Println("Préstamo sin interés")
		}
	default:
		fmt.Println("Sin préstamo otorgado")
	}
}

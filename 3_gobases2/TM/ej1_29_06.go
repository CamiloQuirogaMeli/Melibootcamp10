package main

import "fmt"

func descuento(salario int) int {
	var res int
	if salario > 50000 {
		res = (17 * salario) / 100
		salario = salario - res
		fmt.Println("salario1", salario)
	}
	if salario > 150000 {
		var res1 = (10 * salario) / 100
		salario = salario - res1
		fmt.Println("salario2", salario)
	}
	return salario
}

func main() {
	fmt.Println("Ingrese su salario")
	var i int
	fmt.Scan(&i)
	fmt.Println("Su salario es: ", descuento(i))
}

package main

import "fmt"

func calculoImpuesto(salario float64) float64 {
	if salario >= 50000 {
		return salario * 0.17
	} else if salario >= 150000 {
		return salario * 0.27
	} else {
		return 0
	}
}

func main() {
	var salario float64
	fmt.Println("Porfavor, introduzca su salario")
	fmt.Scanln(&salario)
	fmt.Println("El impuesto al salario es de", calculoImpuesto(salario))
}

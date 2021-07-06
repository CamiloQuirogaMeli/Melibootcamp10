package main

import (
	"fmt"
)

func main() {
	var (
		age, workTime int
		isEmployed    string
		salary        float64
	)
	fmt.Print("Ingrese la edad del cliente: ")
	fmt.Scanln(&age)
	if age > 22 {
		fmt.Print("¿El cliente está empleado? (S/N): ")
		fmt.Scanln(&isEmployed)
		if isEmployed == "S" || isEmployed == "s" {
			fmt.Print("Ingrese la antigüedad del cliente (en años): ")
			fmt.Scanln(&workTime)
			if workTime >= 1 {
				fmt.Print("Ingrese el sueldo del cliente: ")
				fmt.Scanln(&salary)
				if salary >= 100000 {
					fmt.Println("Felicidades, podemos otorgarle el crédito sin interés")
				} else {
					fmt.Println("Felicidades, podemos otorgarle el crédito con intereses bajos")
				}
				return
			}
		}
	}
	fmt.Println("Lo sentimos, no se le puede otorgar un crédito")
}

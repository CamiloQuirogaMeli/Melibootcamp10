package main

import (
	"fmt"
)

func prest() {
	var age int
	var employee string
	var seniority int
	var salary float64

	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&age)
	fmt.Println("Usted se encuentra empleado? Ingrese Y (si) o N (no)")
	fmt.Scanln(&employee)
	fmt.Print("Ingrese su antiguedad en su trabajo ")
	fmt.Scanln(&seniority)
	fmt.Print("Ingrese su sueldo actual ")
	fmt.Scanln(&salary)

	if age > 22 && employee == "Y" && seniority > 1 {
		if salary > 100000 {
			fmt.Println("Se le otorgará un préstamo sin interés")
		} else {
			fmt.Println("Se le otorgará un préstamo pero con interés")
		}
	} else {
		fmt.Println("No es posible otorgarle un préstamo")
	}
}

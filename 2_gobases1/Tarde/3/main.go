package main

import "fmt"

func main() {

	age := 18
	isWorking := true
	antiguatyYears := 2
	salary := 50000

	if age < 23 {
		fmt.Println("Lo sentimos debe ser mayor de 22 años para acceder un préstamo")
	} else if !isWorking {
		fmt.Println("Lo sentimos, debe ser emepleado para acceder a un préstamo")
	} else if antiguatyYears < 2 {
		fmt.Println("Lo sentimos, debe tener al menos dos años de antigüedad para acceder a un préstamo")
	}

	fmt.Println("Felicidades, usted puede acceder a un préstamo")

	if salary > 100000 {
		fmt.Println("Debido al valor de su salario se le cobrará interés")
	}

}

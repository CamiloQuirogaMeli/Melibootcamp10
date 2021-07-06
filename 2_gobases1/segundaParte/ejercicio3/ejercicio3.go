package main

import "fmt"

func main() {
	var age, antiquity, salary int
	var jobOption int
	var employed bool
	errorMessage := "Lo siento, no puedes acceder al prestamo"

	fmt.Println("Escribe tu edad: ")
	fmt.Scanln(&age)
	fmt.Println("Tienes empleo? Escribe 1 para si, 0 para no")
	fmt.Scanln(&jobOption)
	if jobOption == 1 {
		employed = true
	} else {
		employed = false
	}
	fmt.Println("Escribe tu antiguedad (en anios) en el trabajo: ")
	fmt.Scanln(&antiquity)
	fmt.Println("Escribe tu salario: ")
	fmt.Scanln(&salary)

	if age >= 22 && employed && antiquity > 1 {
		if salary > 100000 {
			fmt.Println("Genial! Por tu salario no te cobraremos intereses. Disfruta del prestamo")
		} else {
			fmt.Println("Prestamo exitoso. Disfruta mucho")
		}
	} else {
		fmt.Println(errorMessage)
	}
}

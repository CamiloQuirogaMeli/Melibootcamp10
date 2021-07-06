package main

import "fmt"

func salario() {

	var salary float64

	fmt.Print("Ingrese sus minutos de trabajo: ")
	var minutes float64
	fmt.Scanln(&minutes)
	fmt.Print("Ingrese su categoría: (A o B o C)")
	var category string
	fmt.Scanln(&category)

	switch category {
	case "A":
		salary = 3000 * (minutes / 60)
		salary = salary + salary*0.5
		fmt.Println("El sueldo queda en: ", salary)
	case "B":
		salary = 1500 * (minutes / 60)
		salary = salary + salary*0.2
		fmt.Println("El sueldo queda en: ", salary)
	case "C":
		salary = 1000 * (minutes / 60)
		fmt.Println("El sueldo queda en: ", salary)
	}
}

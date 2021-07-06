/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
minutos trabajados por mes y la categoría, y nos devuelva su salario
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Por favor ingrese los minutos que trabajo sin comas(,) ni puntos(.) -> ")
	var minutes int
	_, errMinutes := fmt.Scanf("%d", &minutes)

	if errMinutes != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Print("Ingrese su categoria (A, B, C)-> ")
	scanner.Scan()
	calcularSalario(minutes, scanner.Text())

}

func calcularSalario(minutes int, category string) {
	if category != "A" && category != "B" && category != "C" {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	var salary int
	switch category {
	case "A":
		salary = categoryA(minutes)
	case "B":
		salary = categoryB(minutes)
	case "C":
		salary = categoryC(minutes)
	}

	println("De acuerdo a los minutos trabajados y a la categoria que pertenece, su salario es de: ", salary)

}

func categoryA(minutes int) (salary int) {
	salary = 3000 / 60
	monthSalary := ((3000 * 24) * 30)
	percent := (monthSalary * 50) / 100
	salary = salary * minutes
	salary = salary + percent
	return
}

func categoryB(minutes int) (salary int) {
	salary = 1500 / 60
	monthSalary := ((1500 * 24) * 30)
	percent := (monthSalary * 20) / 100
	salary = salary * minutes
	salary = salary + percent
	return
}

func categoryC(minutes int) (salary int) {
	salary = 1000 / 60

	salary = salary * minutes
	return
}

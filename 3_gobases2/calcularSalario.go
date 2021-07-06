package main

import "fmt"

//Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que se le pueda pasar como parámetros la cantidad de minutos trabajados por mes y la categoría, y nos devuelva su salario

func main()  {
	fmt.Println("Vamos a calcular el salario para el empleado 1, 2 y 3 con las categorias A, B y C respectivamente.")

	fmt.Println("El salario del empleado 1, que trabajo 90 minutos es $", salario(90, "A"))
	fmt.Println("El salario del empleado 2, que trabajo 40 minutos es $", salario(40, "B"))
	fmt.Println("El salario del empleado 3, que trabajo 120 minutos es $", salario(120, "C"))
}

func salario(min int, cat string) float64 {
	var sal, hora float64

	switch cat {
	case "A":
		hora = 3000 / 60
		sal = hora * float64(min)
		return sal + (sal * 0.5)
	case "B":
		hora = 1500 / 60
		sal = hora * float64(min)
		return sal + (sal * 0.2)
	case "C":
		hora = 1000 / 60
		return hora * float64(min)
	}

	return 0
}

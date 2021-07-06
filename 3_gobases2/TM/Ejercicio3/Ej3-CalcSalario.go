package main

import "fmt"

func calcSalario (minTrabajados float32, cat string) (salarioFinal float32){

	switch cat {
	case "C":
		salarioFinal = (minTrabajados/60)*1000
		return
	case "B":
		salarioFinal = (minTrabajados/60)*1500
		salarioFinal += (salarioFinal*20)/100
		return
	case "A":
		salarioFinal = (minTrabajados/60)*3000
		salarioFinal += (salarioFinal*50)/100
		return
	}

	return
}

func main(){

	var (
		categoria = "C"
		minutosT = float32(43200.5)
	)

	operacion := calcSalario(minutosT, categoria)

	fmt.Printf("El salario a cobrar el empelado es de: $%0.2f", operacion)

}


/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
minutos trabajados por mes y la categoría, y nos devuelva su salario*/
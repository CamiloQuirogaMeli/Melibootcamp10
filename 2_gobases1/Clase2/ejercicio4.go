package main

import (
	"fmt"
)

func main() {

	/*
		Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
		que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
		por qué?
	*/
	/*
		Se puede resolver con condicionales anidados o con condicionales múltiples
		y para guardar los datos se pueden usar maps o listas con un orden específico
	*/

	var numero int
	meses := [12]string{"Enero", "febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "septiembre", "Octubre", "Noviembre", "diciembre"}

	fmt.Println("Digita el número del mes: ")
	fmt.Scanln(&numero)

	fmt.Println("El mes que corresponde es", meses[numero])
}

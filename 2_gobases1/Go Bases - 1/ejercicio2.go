package main

import "fmt"

/*
	Ejercicio 2 - Clima
	Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura
	y humedad y presión atmosférica de distintos lugares.Declara 3 variables especificando el
	tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
	Imprime los valores de las variables en consola.
	¿Qué tipo de dato le asignarías a las variables?
*/

func main() {
	var temperatura, humedad, presion float64 = 18.2, 56, 1027.1
	fmt.Println("La temperatura en Bogotá es: ", temperatura)
	fmt.Println("La humedad en Bogotá es: ", humedad)
	fmt.Println("La presion en Bogotá es: ", presion)
}

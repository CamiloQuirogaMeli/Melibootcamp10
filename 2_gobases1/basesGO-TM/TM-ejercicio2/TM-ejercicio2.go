package main

import "fmt"

/*
Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura
y humedad y presión atmosférica de distintos lugares.
1. Declara 3 variables especificando el tipo de dato, como valor deben tener la
temperatura, humedad y presión de donde te encuentres.
2. Imprime los valores de las variables en consola.
3. ¿Qué tipo de dato le asignarías a las variables?
*/

func main() {

	temperatura := 9.0

	humedad := 60

	presion := 1018.8

	/* Utilice flotantes para temperatura y presion debido a que pueden tomar valor decimal
	Utilice enteros para humedad dado que solo toma valor entero */

	fmt.Println("Datos actuales en Ciudad de Santa Fe\nTemperatura: ",
		temperatura, "ºC\nHumedad: ", humedad, "%\nPresion atmosferica: ",
		presion, "hPa")

}

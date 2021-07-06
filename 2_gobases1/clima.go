package main

import "fmt"

//Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
//1. Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
//2. Imprime los valores de las variables en consola.
//3. ¿Qué tipo de dato le asignarías a las variables? --> int

func main()  {
	var temp, hum, pre int
	temp = 17
	hum = 77
	pre = 1017

	fmt.Println("Información meteorológica de Bogotá, Colombia")
	fmt.Println("Temperatura: ", temp, "°C")
	fmt.Println("Humedad: ", hum, "%")
	fmt.Println("Presión : ", pre, "°hPa")
}
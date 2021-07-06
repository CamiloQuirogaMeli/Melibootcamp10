package main //Define el paquete main

// Importa librerías
import (
	"fmt"
)

func main() {

	/*
		Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura
		y humedad y presión atmosférica de distintos lugares.
		1. Declara 3 variables especificando el tipo de dato, como valor deben tener la
		temperatura, humedad y presión de donde te encuentres.
		2. Imprime los valores de las variables en consola.
		3. ¿Qué tipo de dato le asignarías a las variables?
	*/

	var temperatura float32 //Medido en Centígrados
	var humedad int16       //Medidos en porcentaje
	var presion float32     //Medidos en atmósferas

	temperatura = 17.5
	humedad = 52
	presion = 1.1

	fmt.Println("Para la ciudad de Bogotá, se tiene los siguentes datos: ")
	fmt.Println("La temperatura es de ", temperatura, " grados centigrados")
	fmt.Println("La humedad es de ", humedad, " porciento")
	fmt.Println("La presion atmosférica es de ", presion, " atmósferas")

}

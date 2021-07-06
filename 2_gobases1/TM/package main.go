package main

import "fmt"

func main() {
	ejercicio1()
	ejercicio2()
}

func ejercicio1() {
	var nombre, apellido string = "Luciano", "Dominino"
	fmt.Println("Ejercicio 1")
	fmt.Println(nombre, apellido, "\n")
}

func ejercicio2() {
	var (
		temperatura = 24.2
		humedad     = 58.1
		presion     = 1003.4
	)

	fmt.Println("Ejercicio 2")
	fmt.Println("Temperatura", temperatura, "ºC")
	fmt.Println("Humedad", humedad, "%")
	fmt.Println("Presion", presion, "mb", "\n")
}

func ejercicio3() {
	var nombre string //empezaba declarando la variable con 1
	var apellido string
	var edad int                     //primero indicaba el tipo antes del nombre de la variable
	apellido = "Dominino"            //la variable arrancaba con 1, no hacia falta declarar de manera corta ya que estaba declarada arriba
	var licenciaConducir bool = true //cambio de formato en la forma que se declaraba la variable
	var estaturaPersona int          //igual que la anterior, cambio de formato en la forma de declarar porque tenia espacios, puede ser int si se toman los cm, de lo contrario deberia ser float
	cantidadDeHijos := 2
}

func ejercicio4() {
	var (
		apellido = "Gomez"
		edad     = 35 //estaba mal, lo declaraba como int y le asignaba un string
		//boolean = false //este esta mal declarado, el valor va sin "", ademas le falta un nombre correcto a la variable
		sueldo = 45857.90 //estaba declarado como string y le asignaba un float
		nombre = "Julián"
	)

}

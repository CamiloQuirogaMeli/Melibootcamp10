package main

import "fmt"

func ej1() {
	var name string = "Barrera Julio Franco Nicolas"
	var address string = "Sarmiento 790"
	fmt.Println("Nombre: \t" + name)
	fmt.Println("Direccion: \t" + address)
}
func ej2() {
	var temp, humidity, pressure float64 = 33.5, 50.0, 90.0
	fmt.Println("Temperatura: \t", temp)
	fmt.Println("Humedad: \t", humidity)
	fmt.Println("Presion: \t", pressure)
}
func main() {
	const SEPARATOR = "******---------******"
	fmt.Println(SEPARATOR)
	fmt.Println("--- Ejercicio 1 ---")
	ej1()
	fmt.Println(SEPARATOR)
	fmt.Println("--- Ejercicio 2 ---")
	ej2()
	fmt.Println(SEPARATOR)
}

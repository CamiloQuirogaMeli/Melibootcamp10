package main

import "fmt"

var nombre, direccion string = "Federico", "Murguiondo 1675"

func main() {
	fmt.Println("---Ejercico 1")
	fmt.Println(nombre, direccion)
	fmt.Println("---Ejercico 2")
	ejercicio2()
	fmt.Println("---Ejercico 3")
	ejercicio3()
	fmt.Println("---Ejercico 4")
	ejercicio4()
	fmt.Println("---Ejercico 5")
	cantidadLetras("fede")
	fmt.Println("---Ejercico 6")
	fmt.Println(aplicarDescuento(1889, 50))
	fmt.Println("---Ejercico 7")
	ejercicio7(9)
	fmt.Println("---Ejercico 8")
	listaEstudiantes()
	fmt.Println("---Ejercico 9")
	ejercicio9()

}

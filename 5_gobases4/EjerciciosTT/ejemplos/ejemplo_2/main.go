package main

import "fmt"

func imprimirMensaje1() {
	fmt.Println("1 - Esta función se ejecuta a pesar de producirse panic")
}

func imprimirMensaje2() {
	fmt.Println("2 - Esta función se ejecuta a pesar de producirse panic")
}

func imprimirMensaje3() {
	fmt.Println("3 - Esta función se ejecuta a pesar de producirse panic")
}

func main() {
	//aplicamos "defer" a la invocación de una función anónima
	defer imprimirMensaje1()
	defer imprimirMensaje2()
	defer imprimirMensaje3()
	//creamos un panic con un mensaje de que se produjo
	panic("se produjo panic!!!")

}

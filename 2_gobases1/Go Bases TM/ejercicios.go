package main

import "fmt"

func main() {
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
}

func ejercicio1() {
	fmt.Println("----Ejercicio 1 -----")
	var apellido, direccion string = "Luna", "juan.luna@mercadolibre.com"
	fmt.Println(apellido)
	fmt.Println(direccion)
	fmt.Println("\n")
}

func ejercicio2() {
	fmt.Println("----Ejercicio 2 -----")
	var temperatura float32 = 30
	var humedad uint16 = 100
	var presion float32 = 1015.6
	fmt.Println("Tmperatura: ", temperatura, "ºC")
	fmt.Println("Humedad: ", humedad, "%")
	fmt.Println("Presión:", presion, "mb")
	fmt.Println("\n")
}

func ejercicio3() {
	fmt.Println("----Ejercicio 3 -----")
	fmt.Println("var 1nombre string: Incorrecta, debe iniciar con letras")
	fmt.Println("Correccion: var nombre string")
	fmt.Println("\n")

	fmt.Println("var int edad: Incorrecta, tipo de datos debe ir al final")
	fmt.Println("Correccion: var edad int")
	fmt.Println("\n")

	fmt.Println("var 1apellido: Incorrecta, debe iniciar con letras y tipo de datos incorrecto.")
	fmt.Println("Correccion: var apellido := \"LUNA\"")
	fmt.Println("\n")

	fmt.Println("var licencia_de_conducir = true: Incorrecta, indicar tipo de datos.")
	fmt.Println("Correccion: var licencia_de_conducir boolean = true")
	fmt.Println("\n")

	fmt.Println("var estatura de la persona int: Incorrecta, contiene espacios.")
	fmt.Println("Correccion: var estaturaDeLaPersona int")
	fmt.Println("\n")
}

func ejercicio4() {
	fmt.Println("----Ejercicio 4 -----")
	fmt.Println("edad int = 35")
	fmt.Println("boolean := \"false\"")
	fmt.Println("var sueldo float32 = 45857.90")
	fmt.Println("\n")
}

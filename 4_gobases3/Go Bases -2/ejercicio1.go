package main

import "fmt"

/*
	Ejercicio 1 - Red social

	Una empresa de redes sociales requiere implementar una estructura usuario con funciones
	que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
	que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
	para las funciones:

	La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
	Y deben implementarse las funciones:
		- cambiar nombre: me permite cambiar el nombre y apellido.
		- cambiar edad: me permite cambiar la edad.
		- cambiar correo: me permite cambiar el correo.
		- cambiar contraseña: me permite cambiar la contraseña.
*/

type usuarios struct {
	Nombre     string
	Apellido   string
	Edad       int
	correo     string
	contraseña string
}

func main() {
	var u usuarios

	u.Nombre = "ander"
	u.Apellido = "torres"
	u.Edad = 23
	u.correo = "correo@gmail.com"
	u.contraseña = "12345"

	fmt.Println("ANTES DE CAMBIAR VALORES, CON LA MISMA REFERENCIA DE MEMORIA")
	fmt.Println(u.Nombre)
	fmt.Println(u.Apellido)
	fmt.Println(u.Edad)
	fmt.Println(u.correo)
	fmt.Println(u.contraseña)

	cambiarNombre(&u.Nombre, &u.Apellido)
	cambiarEdad(&u.Edad)
	cambiarCorreo(&u.correo)
	cambiarContraseña(&u.contraseña)

	fmt.Println("")

	fmt.Println("DESPUES DE CAMBIAR VALORES, CON LA MISMA REFERENCIA DE MEMORIA")
	fmt.Println(u.Nombre)
	fmt.Println(u.Apellido)
	fmt.Println(u.Edad)
	fmt.Println(u.correo)
	fmt.Println(u.contraseña)
}

func cambiarNombre(nombre *string, apellido *string) {
	*nombre = "camilo"
	*apellido = "rodriguez"
}

func cambiarEdad(edad *int) {
	*edad = 25
}

func cambiarCorreo(correo *string) {
	*correo = "camilo@gmail.com"
}

func cambiarContraseña(contraseña *string) {
	*contraseña = "54321"
}

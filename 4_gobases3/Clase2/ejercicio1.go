package main

import (
	"fmt"
)

type Usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func (u *Usuario) CambiarNombre(nombre string, apellido string) *Usuario {
	//Devuelve el mismo espacio de memoria
	u.nombre = nombre
	u.apellido = apellido

	fmt.Println("Nombre y apellido cambiado exitosamente")
	return u
}

func (u *Usuario) CambiarEdad(edad int) *Usuario {
	u.edad = edad
	fmt.Println("Edad cambiada exitosamente")
	return u
}

func (u *Usuario) CambiarCorreo(correo string) *Usuario {
	u.correo = correo
	fmt.Println("Correo cambiado exitosamente")
	return u
}

func (u *Usuario) CambiarContrasena(contraseña string) *Usuario {
	u.contraseña = contraseña
	fmt.Println("contraseña cambiada exitosamente")
	return u
}

func main() {
	/*

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
	var nombre string
	var apellido string
	var edad int
	var correo string
	var contrasena string

	usuario1 := Usuario{}

	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")

		fmt.Println("\t 1: Cambiar el nombre y apellido del usuario")
		fmt.Println("\t 2: Cambiar la edad del usuario")
		fmt.Println("\t 3: Cambiar el correo del usuario")
		fmt.Println("\t 4: Cambiar la contraseña del usuario")
		fmt.Println("\t 5: Imprimir Datos del usuario")
		fmt.Println("\t 6: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el nombre del usuario")
			fmt.Scanln(&nombre)
			fmt.Println("Digita el apellido del usuario")
			fmt.Scanln(&apellido)
			usuario1.CambiarNombre(nombre, apellido)
			//fmt.Println("Apuntador desde el main: ", dirUsuario)

		case 2:
			fmt.Println("Digita la edad del usuario")
			fmt.Scanln(&edad)
			usuario1.CambiarEdad(edad)
			//fmt.Println("Apuntador desde el main: ", dirUsuario)

		case 3:
			fmt.Println("Digita el correo del usuario")
			fmt.Scanln(&correo)
			usuario1.CambiarCorreo(correo)
			//fmt.Println("Apuntador desde el main: ", dirUsuario)

		case 4:
			fmt.Println("Digita la contraseña del usuario")
			fmt.Scanln(&contrasena)
			usuario1.CambiarContrasena(contrasena)
			//fmt.Println("Apuntador desde el main: ", dirUsuario)
		case 5:
			fmt.Println("El nombre del usuario es:", usuario1.nombre)
			fmt.Println("El apellido del usuario es:", usuario1.apellido)
			fmt.Println("La edad del usuario es:", usuario1.edad)
			fmt.Println("El correo del usuario es:", usuario1.correo)
		case 6:
			salir = true
		}
	}
}

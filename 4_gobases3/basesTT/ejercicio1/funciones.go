package main

import "fmt"

func cambiarNombre(user *Usuario) {
	var nombreNuevo string
	fmt.Println("Ingrese el nuevo nombre para\t", user.nombre)
	fmt.Scanln(&nombreNuevo)
	user.nombre = nombreNuevo
	fmt.Println("Nombre cambiado\n")
}

func cambiarEdad(user *Usuario) {
	var edad int
	fmt.Println("Ingrese la nueva edad para\t", user.nombre)
	fmt.Scanln(&edad)
	user.edad = edad
	fmt.Println("Edad cambiada\n")
}

func cambiarCorreo(user *Usuario) {
	var correo string
	fmt.Println("Ingrese el nuevo correo para\t", user.nombre)
	fmt.Scanln(&correo)
	user.correo = correo
	fmt.Println("Correo cambiado\n")
}

func cambiarContrasena(user *Usuario) {
	var contra string
	fmt.Println("Ingrese el nuevo pass para\t", user.nombre)
	fmt.Scanln(&contra)
	user.contrasena = contra
	fmt.Println("Password cambiado\n")
}

func imprimirDatos(user *Usuario) {
	fmt.Println("")
	fmt.Println("Nombre ", user.nombre)
	fmt.Println("Apellido ", user.apellido)
	fmt.Println("Edad ", user.edad)
	fmt.Println("Correo ", user.correo)
	fmt.Println("Password ", user.contrasena)
	fmt.Println("")
}

package main

import "fmt"

type Usuario struct {
	nombre   string
	apellido string
	edad     int
	correo   string
	pass     string
}

func cambiarNombre(u *Usuario) {
	var nombre, apellido string
	fmt.Printf("\n***** Cambiar nombre *****\nNuevo nombre: ")
	fmt.Scanf("%s", &nombre)
	(*u).nombre = nombre
	fmt.Printf("Nuevo apellido: ")
	fmt.Scanf("%s", &apellido)
	(*u).apellido = apellido
}

func cambiarEdad(u *Usuario) {
	var edad int
	fmt.Printf("\n***** Cambiar edad *****\nEdad: ")
	fmt.Scanf("%d", &edad)
	(*u).edad = edad
}

func cambiarCorreo(u *Usuario) {
	var correo string
	fmt.Printf("\n***** Cambiar correo *****\nCorreo: ")
	fmt.Scanf("%s", &correo)
	(*u).correo = correo
}

func cambiarPass(u *Usuario) {
	var auxPass string
	fmt.Printf("\n***** Cambiar contraseña *****\nContraseña anterior: ")
	fmt.Scanf("%s", &auxPass)
	if auxPass == u.pass {
		fmt.Printf("Nueva contraseña: ")
		fmt.Scanf("%s", &auxPass)
		(*u).pass = auxPass
		fmt.Printf("Contraseña cambiada con éxito\n\n")
	} else {
		fmt.Printf("Contraseña incorrecta\n\n")
	}
}

func (u Usuario) mostrar() {
	fmt.Printf("Nombre: %s\n", u.nombre)
	fmt.Printf("Apellido: %s\n", u.apellido)
	fmt.Printf("Edad: %d\n", u.edad)
	fmt.Printf("Correo: %s\n", u.correo)
	fmt.Printf("Contraseña: %s\n", u.pass)
}

func main() {
	nuevoUsuario := Usuario{
		nombre:   "Nina",
		apellido: "Samartin",
		edad:     20,
		correo:   "nina.samartin@mercadolibre.com",
		pass:     "123123123",
	}
	nuevoUsuario.mostrar()
	cambiarNombre(&nuevoUsuario)
	cambiarEdad(&nuevoUsuario)
	cambiarCorreo(&nuevoUsuario)
	cambiarPass(&nuevoUsuario)
	nuevoUsuario.mostrar()
}

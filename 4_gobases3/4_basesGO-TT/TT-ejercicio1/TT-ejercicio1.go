package main

import "fmt"

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

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func cambiarNombre(nombreActualizado, apellidoActualizado string, u *usuario) {
	(*u).nombre = nombreActualizado
	(*u).apellido = apellidoActualizado
}

func cambiarEdad(edadActualizada int, u *usuario) {
	(*u).edad = edadActualizada

}

func cambiaCorreo(correoActualizado string, u *usuario) {
	(*u).correo = correoActualizado
}

func cambiarContraseña(contraseñaActualizada string, u *usuario) {
	(*u).contraseña = contraseñaActualizada
}

func listaUsuario(u *usuario) {

	fmt.Println("Nombre: ", (*u).nombre)
	fmt.Println("Apellido: ", (*u).apellido)
	fmt.Println("Edad: ", (*u).edad)
	fmt.Println("Correo: ", (*u).correo)
	fmt.Println("Contraseña: ", (*u).contraseña)

}

func main() {

	/*Genero la estructura*/
	var usuarios []usuario

	u1 := usuario{nombre: "Nombre1", apellido: "Apellido1", edad: 28, correo: "nombreapellido1@gmail.com", contraseña: "1234abc"}
	u2 := usuario{nombre: "Nombre2", apellido: "Apellido2", edad: 27, correo: "nombreapellido2@gmail.com", contraseña: "1234abc"}
	u3 := usuario{nombre: "Nombre3", apellido: "Apellido3", edad: 20, correo: "nombreapellido3@gmail.com", contraseña: "1234abc"}
	u4 := usuario{nombre: "Nombre4", apellido: "Apellido4", edad: 24, correo: "nombreapellido4@gmail.com", contraseña: "1234abc"}

	usuarios = append(usuarios, u1, u2, u3, u4)

	/*Declaro un puntero de tipo usuario para pasar por parámetros*/
	u := &usuarios[0]
	listaUsuario(u)

	cambiarNombre("Modificado", "Modificado", u)
	cambiarEdad(29, u)
	cambiarContraseña("Modificado", u)
	cambiaCorreo("Modificado", u)

	listaUsuario(u)

}

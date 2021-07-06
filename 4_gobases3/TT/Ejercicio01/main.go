package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func cambiarNombre(nombre string, apellido string, usuario *Usuario) {
	usuario.Nombre = nombre
	usuario.Apellido = apellido
}

func cambiarEdad(edad int, usuario *Usuario) {
	usuario.Edad = edad
}

func cambiarCorreo(correo string, usuario *Usuario) {
	usuario.Correo = correo
}

func cambiarContrasena(contrasena string, usuario *Usuario) {
	usuario.Contrasena = contrasena
}

func main() {
	p1 := Usuario{"Marcos", "Gutierrez", 20, "test@gmail.com", "123456"}
	cambiarNombre("Ramon", "Perez", &p1)
	cambiarEdad(30, &p1)
	cambiarCorreo("nuevo.email@mgmail.com", &p1)
	cambiarContrasena("nuevoPass", &p1)
	fmt.Println(p1)
}

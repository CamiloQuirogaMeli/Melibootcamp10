package main

import "fmt"

type Persona struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func cambiarNombre(nombre *string, apellido *string, per *Persona) {
	per.nombre = *nombre
	per.apellido = *apellido
}

func cambiarEdad(edad *int, per *Persona) {
	per.edad = *edad
}

func cambiarCorreo(correo *string, per *Persona) {
	per.correo = *correo
}

func cambiarContraseña(contraseña *string, per *Persona) {
	per.contraseña = *contraseña
}

func main() {

	var (
		nombre     = "Vincent"
		apellido   = "Ramon"
		edad       = 23
		correo     = "vincentconace@gmail.com"
		contraseña = "12345435"
	)

	p1 := Persona{
		nombre:     nombre,
		apellido:   apellido,
		edad:       edad,
		correo:     correo,
		contraseña: contraseña,
	}

	cambiarNombre(&nombre, &apellido, &p1)
	cambiarEdad(&edad, &p1)
	cambiarCorreo(&correo, &p1)
	cambiarContraseña(&contraseña, &p1)

	fmt.Println(p1)
}

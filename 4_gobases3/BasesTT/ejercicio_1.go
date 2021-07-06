package main

import (
	"fmt"
)

type Usuario struct {
	nombre   string
	apellido string
	edad     int
	correo   string
	pass     string
}

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) SetApellido(apellido string) {
	u.apellido = apellido
}

func (u *Usuario) SetEdad(edad int) {
	u.edad = edad
}

func (u *Usuario) SetCorreo(correo string) {
	u.correo = correo
}

func (u *Usuario) SetPass(pass string) {
	u.pass = pass
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetApellido() string {
	return u.apellido
}

func (u *Usuario) GetEdad() int {
	return u.edad
}

func (u *Usuario) GetCorreo() string {
	return u.correo
}

func (u *Usuario) GetPass() string {
	return u.pass
}

func main() {
	usuario := Usuario{}
	fmt.Println("Datos iniciales:", usuario)
	usuario.SetNombre("Luis")
	usuario.SetApellido("Oropeza")
	usuario.SetEdad(26)
	usuario.SetCorreo("luis.oropeza@mercadolibre.com.mx")
	usuario.SetPass("incorrecta")
	fmt.Println("Datos actualizados:", usuario)
}

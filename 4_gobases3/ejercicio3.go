package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       float64
	Correo     string
	Contraseña string
}

func (u *Usuario) setNombre(nombre string) {
	u.Nombre = nombre
}

func (u *Usuario) setApellido(apellido string) {
	u.Apellido = apellido
}

func (u *Usuario) setEdad(edad float64) {
	u.Edad = edad
}

func (u *Usuario) setCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) setContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func ejercicio3() {
	nuevoUsuario := Usuario{
		Nombre:     "fede",
		Apellido:   "Martinez",
		Edad:       27,
		Correo:     "fedemartinez@gmail.com",
		Contraseña: "123456",
	}
	fmt.Println(nuevoUsuario)

	nuevoUsuario.setNombre("Raquel")
	nuevoUsuario.setApellido("Davalos")
	nuevoUsuario.setContraseña("456123")
	nuevoUsuario.setCorreo("Davalosraquel@gmail.com")
	nuevoUsuario.setEdad(31)

	fmt.Println(nuevoUsuario)
}

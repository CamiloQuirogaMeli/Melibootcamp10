package main

import "fmt"

func main() {
	u := usuario{
		Nombre:     "Luciano",
		Apellido:   "Dominino",
		Edad:       "26",
		Correo:     "luciano.dominino@mercadolibre.com",
		Contrasena: "hola123",
	}

	fmt.Println("Usuario original:", u)

	u.cambiarNombre("Lucas")
	fmt.Println("Cambio de nombre:", u)
	fmt.Println("====================")
	u.cambiarEdad("24")
	fmt.Println("Cambio de edad:", u)
	fmt.Println("====================")
	u.cambiarCorreo("lucas.dominino@mercadolibre.com")
	fmt.Println("Cambio de correo:", u)
	fmt.Println("====================")
	u.cambiarContrasena("h123")
	fmt.Println("cambio de contraseña:", u)
	fmt.Println("====================")

}

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       string
	Correo     string
	Contrasena string
}

func (u *usuario) cambiarNombre(nombre string) usuario {
	u.Nombre = nombre
	return *u
}
func (u *usuario) cambiarEdad(edad string) usuario {
	u.Edad = edad
	return *u
}
func (u *usuario) cambiarCorreo(correo string) usuario {
	u.Correo = correo
	return *u
}
func (u *usuario) cambiarContrasena(contrasena string) usuario {
	u.Contrasena = contrasena
	return *u
}

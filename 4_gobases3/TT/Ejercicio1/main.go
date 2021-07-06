package main

import "fmt"

type Usuario struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func (u *Usuario) cambiarNombre(nom string, ape string) {
	u.Nombre = nom
	u.Apellido = ape
}

func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) cambiarContrasenia(pass string) {
	u.Contrasenia = pass
}

func main() {
	u1 := Usuario{}

	fmt.Print(u1)
	u1.cambiarNombre("Juan", "Perez")
	u1.cambiarEdad(35)
	u1.cambiarCorreo("juancito@yahoo.com")
	u1.cambiarContrasenia("12312312312321")

	fmt.Print(u1)

	u1.cambiarNombre("Juan", "Perez Gonzalez")
	fmt.Print(u1)
}

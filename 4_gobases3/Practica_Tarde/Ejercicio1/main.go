package main

import "fmt"

type Usuario struct {
	usuarios    []Usuario
	nombre      string
	apellido    string
	edad        int
	correo      string
	contrasenia string
}

func (u *Usuario) agregarUsuario(usu ...Usuario) {
	u.usuarios = append(u.usuarios, usu...)
}

func (u *Usuario) cambiarNombre(nom, ape string) {
	u.nombre = nom
	u.apellido = ape
}

func (u *Usuario) cambiarEdad(e int) {
	u.edad = e
}

func (u *Usuario) cambiarCorreo(c string) {
	u.correo = c
}

func (u *Usuario) cambiarContra(cont string) {
	u.contrasenia = cont
}

func main() {
	u1 := Usuario{nombre: "Valeria", apellido: "Macina", edad: 33, correo: "valeria.macina@mercadolibre.com", contrasenia: "******"}
	u1.agregarUsuario(u1)
	fmt.Printf(" Nombre: %s,\n Apellido: %s,\n Edad: %d,\n Correo: %s,\n Contrasenia: %s\n", u1.nombre, u1.apellido, u1.edad, u1.correo, u1.contrasenia)
	u1.cambiarNombre("Valentina", "MACINA")
	u1.cambiarEdad(30)
	u1.cambiarCorreo("No tiene")
	u1.cambiarContra("*")
	fmt.Print("-------------------\n ")
	fmt.Printf("Nombre: %s,\n Apellido: %s,\n Edad: %d,\n Correo: %s,\n Contrasenia: %s\n", u1.nombre, u1.apellido, u1.edad, u1.correo, u1.contrasenia)

}

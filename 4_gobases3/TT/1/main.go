package main

import "fmt"

func main() {

	us := usuario{
		nombre:     "Lucas",
		edad:       24,
		correo:     "example@gmail.com",
		contraseña: "clave123",
	}

	fmt.Println(us)

	us.setNombre("Andres Ficarra")
	us.setEdad(23)
	us.setCorreo("example@meli.com")
	us.setClave("nuevaClave123")

	fmt.Print(us)
}

type usuario struct {
	nombre     string
	edad       int
	correo     string
	contraseña string
}

func (u *usuario) setNombre(nombre string) {
	u.nombre = nombre
}

func (u *usuario) setEdad(edad int) {
	u.edad = edad
}
func (u *usuario) setCorreo(correo string) {
	u.correo = correo
}
func (u *usuario) setClave(clave string) {
	u.contraseña = clave
}

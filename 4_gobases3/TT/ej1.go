package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (m *Usuario) SetName(n string, a string) {
	m.Nombre = n
	m.Apellido = a
}
func (m *Usuario) SetEdad(e int) {
	m.Edad = e
}

func (m *Usuario) SetCorreo(cor string) {
	m.Correo = cor
}

func (m *Usuario) SetContraseña(con string) {
	m.Contraseña = con
}
func main() {
	m := Usuario{}
	m.SetName("Daniel", "Sanchez")
	m.SetEdad(26)
	m.SetCorreo("dani@gmail.com")
	m.SetContraseña("***")
	fmt.Println(m)
}

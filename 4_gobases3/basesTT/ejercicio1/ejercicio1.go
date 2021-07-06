package main

import "fmt"

type Usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contrasena string
}

func main() {
	var ptrUser *Usuario
	user := Usuario{
		nombre:     "Franco",
		apellido:   "Ballet",
		edad:       26,
		correo:     "franco.ballet@mercadolibre.com",
		contrasena: "123",
	}
	ptrUser = &user
	imprimirDatos(ptrUser)
	cambiarNombre(ptrUser)
	cambiarEdad(ptrUser)
	cambiarCorreo(ptrUser)
	cambiarContrasena(ptrUser)
	fmt.Println("Usuario actualizado")
	imprimirDatos(ptrUser)
}

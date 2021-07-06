package main

import (
	"fmt"
	"os"
)

type Usuario struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func (u *Usuario) cambiarNombre(valor string) {
	fmt.Println("Nombre anterior: ", u.Nombre)
	u.Nombre = valor
	fmt.Println("Nuevo nombre: ", u.Nombre)
}
func (u *Usuario) cambiarApellido(valor string) {
	fmt.Println("Apellido anterior: ", u.Apellido)
	u.Apellido = valor
	fmt.Println("Nuevo apellido: ", u.Apellido)
}
func (u *Usuario) cambiarEdad(valor int) {
	fmt.Println("Edad anterior: ", u.Edad)
	u.Edad = valor
	fmt.Println("Nueva edad: ", u.Edad)
}
func (u *Usuario) cambiarCorreo(valor string) {
	fmt.Println("Correo anterior: ", u.Correo)
	u.Correo = valor
	fmt.Println("Nuevo correo: ", u.Correo)
}
func (u *Usuario) cambiarContrasenia(valor string) {
	fmt.Println("Contrasenia anterior: ", u.Contrasenia)
	u.Contrasenia = valor
	fmt.Println("Nueva contrasenia: ", u.Contrasenia)
}

func main() {
	usuario := Usuario{}
	opcion := 0
	for opcion != 6 {
		fmt.Println("1. Cambiar nombre")
		fmt.Println("2. Cambiar apellido")
		fmt.Println("3. Cambiar edad")
		fmt.Println("4. Cambiar correo")
		fmt.Println("5. Cambiar contrasenia")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			nom := ""
			fmt.Println("Ingrese el nuevo nombre: ")
			fmt.Scanln(&nom)
			usuario.cambiarNombre(nom)
		case 2:
			ape := ""
			fmt.Println("Ingrese el nuevo apellido: ")
			fmt.Scanln(&ape)
			usuario.cambiarApellido(ape)
		case 3:
			ed := 0
			fmt.Println("Ingrese nueva edad: ")
			fmt.Scanln(&ed)
			usuario.cambiarEdad(ed)
		case 4:
			cor := ""
			fmt.Println("Ingrese el nuevo correo: ")
			fmt.Scanln(&cor)
			usuario.cambiarCorreo(cor)
		case 5:
			cont := ""
			fmt.Println("Ingrese la nueva contrasenia: ")
			fmt.Scanln(&cont)
			usuario.cambiarContrasenia(cont)
		case 6:
			os.Exit(1)
		}
	}
}

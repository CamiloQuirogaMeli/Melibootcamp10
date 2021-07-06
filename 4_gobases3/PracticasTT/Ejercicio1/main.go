package main

import "fmt"

type usuarios struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func main() {
	usuario := usuarios{
		nombre:     "Lukas",
		apellido:   "moreno",
		edad:       23,
		correo:     "lukas@hotmail.com",
		contraseña: "54321",
	}
	nombre := "Camilo"
	apellido := "Quiroga"
	edad := 22
	correo := "camilo@gmail.com"
	contraseña := "1234"
	//usuario.nombre = nombre
	//usuario.apellido = apellido
	//usuario.edad = edad
	//usuario.correo = correo
	//usuario.contraseña = contraseña

	fmt.Println(usuario)
	usuario.cambiarnombre(&nombre, &apellido)
	usuario.cambiaredad(&edad)
	usuario.cambiarcorreo(&correo)
	usuario.cambiarcontraseña(&contraseña)
	fmt.Println(usuario)

}

func (nombre *usuarios) cambiaredad(edad *int) {
	nombre.edad = *edad
}
func (nombre *usuarios) cambiarcorreo(correo *string) {
	nombre.correo = *correo
}
func (nombre *usuarios) cambiarcontraseña(contraseña *string) {
	nombre.contraseña = *contraseña
}

func (nombre *usuarios) cambiarnombre(cnombre *string, capellido *string) {

	nombre.nombre = *cnombre
	nombre.apellido = *capellido

}

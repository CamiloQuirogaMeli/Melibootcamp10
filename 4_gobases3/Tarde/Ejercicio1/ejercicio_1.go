package main

type usuario struct {
	nombre     string
	apellido   string
	edad       uint
	correo     string
	contraseña string
}

func (u *usuario) cambiarCorreo(c string) {
	u.correo = c
}

func (u *usuario) cambiarNombre(n string) {
	u.nombre = n
}

func (u *usuario) cambiarEdad(e uint) {
	u.edad = e
}

func (u *usuario) cambiarContraseña(c string) {
	u.contraseña = c
}

func main() {
	u := usuario{"Adrian", "Nanin", 24, "adrian.nanin@mercadolibre.com", "123"}

	u.cambiarEdad(25)
}

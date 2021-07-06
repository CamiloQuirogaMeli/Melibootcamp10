package main

type Usuario struct {
	Nombre     string
	Apellido   string
	edad       int
	correo     string
	contrasena string
}

func (u *Usuario) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int) {
	u.edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {
	u.correo = correo
}

func (u *Usuario) cambiarContrasena(contrasena string) {
	u.contrasena = contrasena
}

func main() {

	usuario := Usuario{}

	usuario.cambiarContrasena("sdsdfcef")
	usuario.cambiarCorreo("ej@ej.com")
	usuario.cambiarEdad(23)
	usuario.cambiarNombre("JuanPablo", "gomez")

}

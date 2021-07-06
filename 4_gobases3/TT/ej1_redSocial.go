package main

/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
- cambiar nombre: me permite cambiar el nombre y apellido.
- cambiar edad: me permite cambiar la edad.
- cambiar correo: me permite cambiar el correo.
- cambiar contraseña: me permite cambiar la contraseña.
*/
import "fmt"

const (
	nombreyapellido = "nombreyapellido"
	edad            = "edad"
	correo          = "correo"
	contrasena      = "contrasena"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (u *Usuario) cambiarNombreyApellido() {
	fmt.Print("Ingrese el nombre nuevo: \n")
	var nombreNuevo string
	fmt.Scanln(&nombreNuevo)
	fmt.Print("Ingrese el apellido nuevo: \n")
	var apellidoNuevo string
	fmt.Scanln(&apellidoNuevo)

	u.Nombre = nombreNuevo
	u.Apellido = apellidoNuevo
	fmt.Println("Usuario modificiado con éxito")
}

func (u *Usuario) cambiarEdad() {
	fmt.Print("Ingrese la edad nueva: \n")
	var edadNueva int
	fmt.Scanln(&edadNueva)

	u.Edad = edadNueva
	fmt.Println("Usuario modificiado con éxito")
}

func (u *Usuario) cambiarCorreo() {
	fmt.Print("Ingrese el correo nuevo: \n")
	var correoNuevo string
	fmt.Scanln(&correoNuevo)

	u.Correo = correoNuevo
}

func (u *Usuario) cambiarContrasenia() {
	fmt.Print("Ingrese la contraseña nueva: \n")
	var contrasenaNueva string
	fmt.Scanln(&contrasenaNueva)

	u.Contrasena = contrasenaNueva
	fmt.Println("Usuario modificiado con éxito")
}

func procesoModificar(u *Usuario, valor string) (usuario string) {
	switch valor {
	case nombreyapellido:
		u.cambiarNombreyApellido()
	case edad:
		u.cambiarEdad()
	case correo:
		u.cambiarCorreo()
	case contrasena:
		u.cambiarContrasenia()
	}
	return
}

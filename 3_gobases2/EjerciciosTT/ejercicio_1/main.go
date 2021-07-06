package main

import "fmt"

type Alumno struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

// Forma 1: definiendo métodos para la estructura
// POO en Go
func (a *Alumno) CambiarNombre(nombre, apellido string) {
	a.Nombre = nombre
	a.Apellido = apellido
}

func (a *Alumno) CambiarEdad(edad int) {
	a.Edad = edad
}

func (a *Alumno) CambiarCorreo(correo string) {
	a.Correo = correo
}

func (a *Alumno) CambiarContraseña(contraseña string) {
	a.Contrasenia = contraseña
}

// Forma dos: pasando referencia a la estructura
// Pasando la estructura como referencia
func CambiarNombreDos(a *Alumno, nombre, apellido *string) {
	a.Nombre = *nombre
	a.Apellido = *apellido
}

func CambiarEdadDos(a *Alumno, edad *int) {
	a.Edad = *edad
}

func CambiarCorreoDos(a *Alumno, correo *string) {
	a.Correo = *correo
}
func CambiarContraseniaDos(a *Alumno, contrasenia *string) {
	a.Contrasenia = *contrasenia
}

func mostrarAlumnos(alumno Alumno) {
	fmt.Println("   Nombre:   ", alumno.Nombre)
	fmt.Println("   Apellido: ", alumno.Apellido)
	fmt.Println("   Edad:      ", alumno.Edad)
	fmt.Println("   Correo:    ", alumno.Correo)
	fmt.Println("   Contraseña:    ", alumno.Contrasenia)
}

func main() {
	// Forma uno
	a := Alumno{
		Nombre:      "Mario",
		Apellido:    "Rosales",
		Edad:        30,
		Correo:      "marioalberto.rosales@mercadolibre.com",
		Contrasenia: "12345",
	}

	mostrarAlumnos(a)

	a.CambiarNombre("Mareque", "Rosales 2")
	a.CambiarEdad(32)
	a.CambiarCorreo("otro.correo@otrodomini.com")
	a.CambiarContraseña("nueva contrasenia")

	mostrarAlumnos(a)

	// otra forma de hacerlo

	a2 := Alumno{}
	var nombre string = "Nuevo Nombre"
	var apellido string = "Nuevo Apellido"
	var edad int = 35
	var correo string = "nuevo.correo@nuevodominio.com"
	var contrasenia string = "contrasenia cualquiera"
	CambiarNombreDos(&a2, &nombre, &apellido)
	CambiarEdadDos(&a2, &edad)
	CambiarCorreoDos(&a2, &correo)
	CambiarContraseniaDos(&a2, &contrasenia)

	mostrarAlumnos(a2)
}

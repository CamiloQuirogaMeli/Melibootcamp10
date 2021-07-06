package main

import f "fmt"

func main() {
	/*
		Ejercicio 1 - Red social
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

	newUser := User{}

	var name string
	var surname string
	var age int
	var email string
	var pass string

	f.Println("Ingresar Nombre")
	f.Scanln(&name)
	f.Println("Ingresar Apellido")
	f.Scanln(&surname)
	f.Println("Ingresar Edad")
	f.Scanln(&age)
	f.Println("Ingresar Email")
	f.Scanln(&email)
	f.Println("Ingresar Contrasena")
	f.Scanln(&pass)

	newUser.changeName(name, surname)
	newUser.changeAge(age)
	newUser.changeEmail(email)
	newUser.changePass(pass)

	newUser.details()
}

type User struct {
	name    string
	surname string
	age     int
	email   string
	pass    string
}

func (u *User) changeName(newName string, newSurname string) {
	u.name = newName
	u.surname = newSurname
}
func (u *User) changeAge(newAge int) {
	u.age = newAge
}
func (u *User) changeEmail(newEmail string) {
	u.email = newEmail
}
func (u *User) changePass(newPass string) {
	u.pass = newPass
}
func (u *User) details() {
	f.Println(
		"Nombre Completo: ", u.name, " ", u.surname, "\n",
		"Edad: ", u.age, "\n",
		"Email: ", u.email, "\n",
		"Contrasena: ", u.pass)
}

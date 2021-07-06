package main

import "fmt"

type user struct {
	name    string
	surname string
	age     int
	email   string
	pass    string
}

func main() {

	newUser := user{"Diana", "Guggeri", 4, "diana.guggeri@gmail.com", "1234ASD"}

	fmt.Printf("%+v\n", newUser)

	changeName(&newUser, "Diani", "Gugg")
	changeAge(&newUser, 4)
	changeEmail(&newUser, "dguggeri@gmail.com")
	changePass(&newUser, "ABC123")

	fmt.Printf("%+v", newUser)
}

func changeName(u *user, newName string, newSurname string) {
	u.name = newName
	u.surname = newSurname
}

func changeAge(u *user, newAge int) {
	u.age = newAge
}

func changeEmail(u *user, newEmail string) {
	u.email = newEmail
}

func changePass(u *user, newPass string) {
	u.pass = newPass
}

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

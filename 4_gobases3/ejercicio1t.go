package main

import "fmt"

/* Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
- cambiar nombre: me permite cambiar el nombre y apellido.
- cambiar edad: me permite cambiar la edad.
- cambiar correo: me permite cambiar el correo.
- cambiar contraseña: me permite cambiar la contraseña. */

type User struct {
	name     string
	lastName string
	age      int
	mail     string
	pass     string
}

func changeName(name, lastName *string, u *User) {
	u.name = name
	u.lastName = lastName
}
func changeAge(age *int, u *User) {
	u.age = age
}
func changeMail(mail *string, u *User) {
	u.mail = mail
}
func changePass(pass *string, u *User) {
	u.pass = pass
}

func (u User) print() {
	fmt.Println("Nombre: ", u.name, " ", u.lastName)
	fmt.Println("Edad: ", u.age)
	fmt.Println("Correo: ", u.mail)
	fmt.Println("Contraseña: ", u.pass)

}

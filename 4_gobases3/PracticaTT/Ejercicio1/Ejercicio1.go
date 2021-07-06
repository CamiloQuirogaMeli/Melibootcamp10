// Ejercicio 1 - Red social
// Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
// Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// Y deben implementarse las funciones:
// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.

package main

import (
	f "fmt"
)

type User struct {
	name     string
	lastname string
	age      int
	mail     string
	password string
}

func (u *User) changeName(name string) {
	u.name = name
}

func (u *User) changeLastName(lastName string) {
	u.lastname = lastName
}

func (u *User) changeAge(age int) {
	u.age = age
}

func (u *User) changePassword(password string) {
	u.password = password
}

func (u *User) changeMail(mail string) {
	u.mail = mail
}

func main() {

	user1 := User{
		name:     "Pedro",
		lastname: "Barrios",
		age:      25,
		mail:     "pedroB@gmail.com",
		password: "123",
	}

	f.Printf("El nombre del usuario es: %s\n", user1.name)
	f.Printf("El apellido del usuario es: %s\n", user1.lastname)
	f.Printf("La edad del usuario es: %d\n", user1.age)
	f.Printf("El mail del usuario es: %s\n", user1.mail)
	f.Printf("La password del usuario es: %s\n", user1.password)

	f.Printf("Ahora se modificaran los datos del usuario\n")

	user1.changeName("Gabriel")
	user1.changeLastName("Sandoval")
	user1.changeAge(22)
	user1.changeMail("gabrielsandoval95@gmail.com")
	user1.changePassword("123456")

	f.Printf("Datos del usuario modificados: \n")

	f.Printf("El nombre del usuario es: %s\n", user1.name)
	f.Printf("El apellido del usuario es: %s\n", user1.lastname)
	f.Printf("La edad del usuario es: %d\n", user1.age)
	f.Printf("El mail del usuario es: %s\n", user1.mail)
	f.Printf("La password del usuario es: %s\n", user1.password)

}

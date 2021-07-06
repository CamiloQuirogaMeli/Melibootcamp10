/* Ejercicio 1
Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
- cambiar nombre: me permite cambiar el nombre y apellido.
- cambiar edad: me permite cambiar la edad.
- cambiar correo: me permite cambiar el correo.
- cambiar contraseña: me permite cambiar la contraseña. */

package main

import (
	"fmt"
)

type Usuario struct {
	Nombre string
	Apellido string
	edad int
	correo string
	contraseña string
}
func (u *Usuario) changeName(nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}
func (u *Usuario) changeAge(edad int) {
	u.edad = edad
}
func (u *Usuario) changeEmail(email string) {
	u.correo = email
}
func (u *Usuario) changePass(pw string) {
	u.contraseña = pw
}
var testUser Usuario = Usuario{"Gonzalo", "Rodriguez", 24, "gonzalodemian.rodriguez@mercadolibre.com :D", "superSecret1234"}

func main() {
	testUser.changeName("Gonzalo Demian", "Rodriguez")
	fmt.Println(testUser)
}
package main

import "fmt"

type User struct {
	Nombre      string
	Edad        int
	Correo      string
	Contrasenia string
}

func main() {
	u1 := User{Nombre: "Joni", Edad: 25, Correo: "jonikrucheski@gmail.com", Contrasenia: "admin"}
	u := &u1
	fmt.Println(u1)

	changeName(u, "Jonathan")
	fmt.Println(*u)

	changeAge(u, 27)
	fmt.Println(u)

	changeEmail(u, "jonathankrucheski@gmail.com")
	fmt.Println(u)

	changePass(u, "1234")
	fmt.Println(u)

}

func changeName(u *User, name string) {
	u.Nombre = name
}

func changeAge(u *User, age int) {
	u.Edad = age
}

func changeEmail(u *User, email string) {
	u.Correo = email
}
func changePass(u *User, pas string) {
	u.Contrasenia = pas
}

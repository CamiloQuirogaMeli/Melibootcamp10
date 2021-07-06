package main

import "fmt"

type User struct {
	name     string
	age      int
	email    string
	password string
}

func setName(u *User, name string) {
	u.name = name
}
func setAge(u *User, age int) {
	u.age = age
}
func setEmail(u *User, email string) {
	u.email = email
}
func setPassword(u *User, password string) {
	u.password = password
}

func main() {
	var option int
	var aux int = 1
	var pointer *User
	user := User{"Random Name", 1, "random@email.com", "randompass"}
	pointer = &user

	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Cambiar nombre\n2. Cambiar edad\n3. Cambiar correo\n4. Cambiar contraseña\n5. Ver usuario \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var name string
			fmt.Println("Ingrese el nuevo nombre")
			fmt.Scan(&name)
			setName(pointer, name)
			fmt.Println("El nombre ha sido actualizado exitosamente")
		case 2:
			var age int
			fmt.Println("Ingrese la nueva edad")
			fmt.Scan(&age)
			setAge(pointer, age)
			fmt.Println("La edad ha sido actualizada exitosamente")
		case 3:
			var email string
			fmt.Println("Ingrese el nuevo correo")
			fmt.Scan(&email)
			setEmail(pointer, email)
			fmt.Println("El correo ha sido actualizado exitosamente")
		case 4:
			var password string
			fmt.Println("Ingrese la nueva contraseña")
			fmt.Scan(&password)
			setPassword(pointer, password)
			fmt.Println("La contraseña ha sido actualizado exitosamente")
		case 5:
			fmt.Println("***** USUARIO *****")
			fmt.Printf("Nombre: %s\nEdad: %d\nCorreo: %s\nPassword: %s\n", user.name, user.age, user.email, user.password)
		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}
}

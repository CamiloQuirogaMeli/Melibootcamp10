package main

import (
	"fmt"
)

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

//falta consultar porqué al modificar no me toma el cambio a imprimirlo

func setNameLastname(user1 *User, uName string, uApellido string) {
	user1.Name = uName
	user1.Lastname = uApellido
}

func setAge(user1 *User, uAge int) {
	user1.Age = uAge
}

func setEmail(user1 *User, uEmail string) {
	user1.Email = uEmail
}

func setPassword(user1 *User, uPassword string) {
	user1.Password = uPassword
}

func main() {
	user1 := User{"Martin", "Hernandez", 35, "mh@gmail.com", "pass"}
	var name, lastname, email, password string
	var age int
	// fmt.Println("Ingrese el nombre: ")
	// fmt.Scanf("%s", &user1.Name)
	// fmt.Println("Ingrese el apellido: ")
	// fmt.Scanf("%s", &user1.Lastname)
	// fmt.Println("Ingrese la edad: ")
	// fmt.Scanf("%d", &user1.Age)
	// fmt.Println("Ingrese el email: ")
	// fmt.Scanf("%s", &user1.Email)
	// fmt.Println("Ingrese el password: ")
	// fmt.Scanf("%s", &user1.Password)
	// fmt.Println(user1)
	var option int = 0
	fmt.Println("Ingrese la opcion deseada: ")
	fmt.Println("1 - Modificar nombre y apellido")
	fmt.Println("2 - Modificar edad")
	fmt.Println("3 - Modificar email")
	fmt.Println("4 - Modificar contraseña")
	//fmt.Scanf("%d", &option)
	switch {
	case option == 0:
		//para hacer debug modifico nombre forzado
		setNameLastname(&user1, "Horacio", "Lawrie")
	case option == 1:
		fmt.Println("Ingrese el nombre: ")
		fmt.Scanf("%s", name)
		fmt.Println("Ingrese el apellido: ")
		fmt.Scanf("%s", lastname)
		setNameLastname(&user1, name, lastname)
	case option == 2:
		fmt.Println("Ingrese la edad: ")
		fmt.Scanf("%d", age)
		setAge(&user1, age)
	case option == 3:
		fmt.Println("Ingrese el email: ")
		fmt.Scanf("%s", email)
		setEmail(&user1, email)
	case option == 4:
		fmt.Println("Ingrese la contraseña: ")
		fmt.Scanf("%s", password)
		setPassword(&user1, password)
	default:
		fmt.Println("No ingresó la opcion correcta")
	}

	fmt.Println(user1)

}

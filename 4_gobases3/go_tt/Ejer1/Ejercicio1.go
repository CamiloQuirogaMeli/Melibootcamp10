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
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       string
	Correo     string
	Contrasena string
}

func cambiarNombre(u *Usuario) {
	scanner := bufio.NewScanner(os.Stdin)
	println("Por favor escriba el nombre")
	scanner.Scan()
	u.Nombre = scanner.Text()
}

func cambiarApellido(u *Usuario) {
	scanner := bufio.NewScanner(os.Stdin)
	println("Por favor escriba el apellido")
	scanner.Scan()
	u.Apellido = scanner.Text()
}

func cambiarCorreo(u *Usuario) {
	scanner := bufio.NewScanner(os.Stdin)
	println("Por favor escriba el correo")
	scanner.Scan()
	u.Correo = scanner.Text()
}

func cambiarContrasena(u *Usuario) {
	scanner := bufio.NewScanner(os.Stdin)
	println("Por favor escriba la contraseña")
	scanner.Scan()
	u.Contrasena = scanner.Text()
}

func main() {
	OPTIONS := `
	1. Cambiar nombre
	2. Cambiar apellido
	3. cambiar correo
	4. cambiar contraseña
	5. Ver datos de usuario
	6. Salir
	`
	user := &Usuario{}
	fmt.Println("Bienvenido al programa")

	continous := true
	for continous {
		fmt.Println("Digite el numero de la funcion que desee realizar")
		fmt.Println(OPTIONS)
		var option int
		_, errOption := fmt.Scanf("%d", &option)

		if errOption != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}

		switch option {
		case 1:
			cambiarNombre(user)
		case 2:
			cambiarApellido(user)
		case 3:
			cambiarCorreo(user)
		case 4:
			cambiarContrasena(user)
		case 5:

			fmt.Println("Datos del usuario", user)
		case 6:
			fmt.Println("Gracias por utilizar el programa")
			os.Exit(0)
		default:
			fmt.Println("La opcion no existe")
			os.Exit(0)
		}
	}

}

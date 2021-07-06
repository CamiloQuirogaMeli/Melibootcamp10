package main

import "fmt"

/*
1. Crea una aplicación donde tengas como variable tu nombre y dirección.
2. Imprime en consola el valor de cada una de las variables.
*/

func main() {

	var nombre string
	nombre = "Agostina Reinhardt"

	var direccion string
	direccion = "Urquiza 2714"

	fmt.Println("Mi nombre es: ", nombre, "\n", "Mi direccion es: ", direccion)

}

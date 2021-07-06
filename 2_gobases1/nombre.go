package main

import "fmt"

//Crea una aplicación donde tengas como variable tu nombre y dirección.
//Imprime en consola el valor de cada una de las variables.

var name, addr = "Juanita", "Calle 66B 70 40"

func main()  {
	fmt.Println("Nombre: " + name)
	fmt.Println("Dirección: " + addr)
}


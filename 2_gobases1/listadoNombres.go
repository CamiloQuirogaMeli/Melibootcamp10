package main

import "fmt"

//a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es necesario generar una aplicación que contenga dicha lista. Estudiantes: Benjamin,
//Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan, Leandro, Eduardo, Duvraschka.

//b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente. Estudiante: Gabriela

func main()  {
	var names = []string {"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(names)
	names = append(names, "Gabriela")
	fmt.Println(names)
}
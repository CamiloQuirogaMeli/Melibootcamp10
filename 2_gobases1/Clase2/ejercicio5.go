package main

import (
	"fmt"
)

func main() {

	/*
		Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
		necesario generar una aplicación que contenga dicha lista.
		Estudiantes:
		Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
		Leandro, Eduardo, Duvraschka.

		b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
		sin modificar el código que escribiste inicialmente.
	*/

	estudiantes := make([]string, 13)
	var estudiante string
	estudiantes[0] = "Benjamin"
	estudiantes[1] = "Nahuel"
	estudiantes[2] = "Brenda"
	estudiantes[3] = "Marcos"
	estudiantes[4] = "Pedro"
	estudiantes[5] = "Axel"
	estudiantes[6] = "Alez"
	estudiantes[7] = "Dolores"
	estudiantes[8] = "Federico"
	estudiantes[9] = "Hernan"
	estudiantes[10] = "Leandro"
	estudiantes[11] = "Eduardo"
	estudiantes[12] = "Duvraschka"

	fmt.Println("Los estudiantes son: ")
	imprimirArreglo(estudiantes, "Estudiante")

	fmt.Println("Digita el nombre del estudiante nuevo (Sugerencia : Gabriela): ")
	fmt.Scanln(&estudiante)
	estudiantes = append(estudiantes, estudiante)
	fmt.Println("Los estudiantes son: ")
	imprimirArreglo(estudiantes, "Estudiante")

}

func imprimirArreglo(arreglo []string, stringItem string) {
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(stringItem, i+1, ":", arreglo[i])
	}
}

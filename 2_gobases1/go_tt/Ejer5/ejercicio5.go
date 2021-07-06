/*
a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
necesario generar una aplicación que contenga dicha lista.
Estudiantes:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
Leandro, Eduardo, Duvraschka.

b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
sin modificar el código que escribiste inicialmente.
Estudiante:
Gabriela
*/
// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	students = append(students, "Gabriela")
	clases := "si"
	for clases == "si" {
		println("Los estudiantes de la clase son: ")
		for i, student := range students {
			fmt.Println(i, student)
		}
		println("¿Desea agregar a otro estudiante? (Y o N)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "N" || scanner.Text() == "n" {
			clases = "Finalizada"
		} else if scanner.Text() == "Y" || scanner.Text() == "y" {
			fmt.Print("Digite el nombre del nuevo estudiante ->")
			scanner.Scan()
			students = append(students, scanner.Text())
		}
	}
}

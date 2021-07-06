package main

import "fmt"

func main() {
	var option int

	const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

	const options string = "1. Ver lista de estudiantes \n 2. Agregar estudiante \n 0. Salir"

	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	var aux int = 1
	for aux != 0 {
		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)
		aux = option
		if option == 1 {
			for _, student := range students {
				fmt.Println(string(student))
			}
		} else if option == 2 {
			var newStudent string
			fmt.Println("Nombre del estudiante")
			fmt.Scan(&newStudent)

			students = append(students, newStudent)
		} else if option == 0 {
			break
		} else {
			fmt.Println("Opcion incorrecta!!!!!")
		}

	}
}

/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.
Por otro lado también es necesario:
	- Saber cuántos de sus empleados son mayores a 21 años.
	- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	- Eliminar a Pedro del mapa.
*/
// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	options := `
	1. Agregar empleado
	2. Modificar Empleado
	3. Eliminar Empleado
	4. Listar empleados mayores a 21 años
	5. Salir del programa
	`
	continous := true
	for continous == true {
		println()
		println("Los empleados son los siguientes: ")
		for employee, age := range employees {
			println(employee, age)
		}
		println()
		println("Seleccione la opción que desea ejecutar")
		println(options)
		print("-> ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			println("Agregar empleado")
			print("Por favor escriba el nombre del nuevo empleado ->")
			scanner.Scan()
			nombre := scanner.Text()
			println()
			println("Digite la edad del empleado en años")
			var age int
			fmt.Print("-> ")
			_, errAge := fmt.Scanf("%d", &age)

			if errAge != nil {
				fmt.Println("Bad Input")
				os.Exit(0)
			}
			employees[nombre] = age
		case "2":
			println("Modificar empleado")
			print("Digite el nombre del empleado el cual desee modificar ->")
			scanner.Scan()
			var nombre = scanner.Text()
			println()
			println("Digite la edad del empleado en años ->")
			fmt.Print("-> ")
			var age int
			_, errAge := fmt.Scanf("%d", &age)

			if errAge != nil {
				fmt.Println("Bad Input")
				os.Exit(0)
			}
			employees[nombre] = age
		case "3":
			println("Eliminar empleado")
			print("Digite el nombre del empleado el cual desee eliminar ->")
			scanner.Scan()
			delete(employees, scanner.Text())
		case "4":
			println("Listado de empleados mayores a 21 años")
			for employee, age := range employees {
				if age > 21 {
					println(employee, age)
				}
			}
		case "5":
			os.Exit(0)
		default:
			os.Exit(0)
		}

	}
}

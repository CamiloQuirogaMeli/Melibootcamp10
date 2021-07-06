package main

import "fmt"

func main() {
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	option:=1
	for option > 0 && option < 5{	
	option:=menu()
	if option != 1 && option != 2 && option != 3 && option != 4 {
		break
	}	
	switch option {

	case 1: 
		var empleado string
		fmt.Println("Ingrese el nombre de un empleado: ")
        	fmt.Scanln(&empleado)
		fmt.Println(empleado, "tiene", empleados[empleado], "años")
	case 2:
		var mayores int
		for _, edad := range empleados {
		if edad >= 21 {
			mayores++
			}
		}
		fmt.Println(mayores, " empleados son mayores de 21 años")
	case 3:
		var empleado string
                fmt.Println("Ingrese el nombre de un empleado: ")
                fmt.Scanln(&empleado)
		var edad int
                fmt.Println("Ingrese la edad del empleado: ")
                fmt.Scanln(&edad)
		//empleados["Federico"] = 25
		empleados[empleado] = edad
		fmt.Println(empleados)
	case 4: 
		var empleado string
                fmt.Println("Ingrese el nombre del empleado que quiere eliminar: ")
                fmt.Scanln(&empleado)
		delete(empleados, empleado)
		fmt.Println(empleados)
		}
	}
	fmt.Println("")
}

func menu()int{
        var option int
        fmt.Println("MENU")
        fmt.Println("1 - Visualizar la edad de un empleado")
        fmt.Println("2 - Visualizar los empleados mayores de 21 años")
        fmt.Println("3 - Agregar un empleado")
        fmt.Println("4 - Eliminar un empleado")
	fmt.Println("Presione otra tecla para salir")
        fmt.Println("Ingrese una opcion: ")
        fmt.Scanln(&option)
	return option
}

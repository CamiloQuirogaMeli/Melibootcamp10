package main

import "fmt"

func main() {
	var (
		edad       = 22
		empleado   = true
		antiguedad = 12 //Expresado en meses
		sueldo     = 100000
	)

	switch {
	case edad < 22:
		fmt.Println("El cliente no cuenta con edad suficiente para acceder al prestamo")
	case !empleado:
		fmt.Println("El beneficio es solo para clientes empleados")
	case antiguedad < 12:
		fmt.Println("El cliente no cuenta con la antiguedad requerida para acceder al prestamo")
	case (edad >= 22 && empleado && antiguedad >= 12 && sueldo < 100000):
		fmt.Println("El cliente puede acceder al beneficio")
	case (edad >= 22 && empleado && antiguedad >= 12 && sueldo >= 100000):
		fmt.Println("Su sueldo lo habilita a un beneficio especial sin intereses")
	default:
		fmt.Println("Verifique los datos ingresados")
	}

	/* 	if edad < 22 {
	   		fmt.Println("El cliente no cuenta con edad suficiente para acceder al prestamo")
	   	} else if !empleado {
	   		fmt.Println("El beneficio es solo para clientes empleados")
	   	} else if antiguedad < 12 {
	   		fmt.Println("El cliente no cuenta con la antiguedad requerida para acceder al prestamo")
	   	} else {
	   		if sueldo < 100000 {
	   			fmt.Println("El cliente puede acceder al beneficio")
	   		} else {
	   			fmt.Println("Su sueldo lo habilita a un beneficio especial sin intereses")
	   		}
	   	}
	*/
}

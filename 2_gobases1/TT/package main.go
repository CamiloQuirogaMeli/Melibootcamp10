package main

import (
	"fmt"
)

func main() {
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
}

func ejercicio1() {
	fmt.Println("Ejercicio 1")
	var palabra string = "Dominino"
	fmt.Println("Palabra elegida:", palabra)
	fmt.Println("Cantidad de letras:", len(palabra))
	fmt.Print("Letras: ")
	for i := 0; i < len(palabra); i++ {
		fmt.Print(string(palabra[i]))
		if i != len(palabra)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("\n")
	fmt.Println("=================")
}

func ejercicio2() {
	fmt.Println("Ejercicio 2")
	var (
		precio    = 1000.00
		descuento = 20
	)
	descuentoCalculado := precio * (float64(descuento) / 100)
	precioConDescuento := precio - descuentoCalculado
	fmt.Println("El valor final del producto con precio $", precio, "con descuento del", descuento, "% es de: $", precioConDescuento)
	fmt.Println("=================")
}

func ejercicio3() {
	fmt.Println("Ejercicio 3")
	var (
		edad       = 23
		antiguedad = 2    //es valor calculado de fecha ingreso con fecha actual
		empleado   = true //establece si el empleado sigue activo (true) o no (false) en la empresa
		sueldo     = 100000.00
	)
	if edad > 22 && antiguedad > 1 && empleado {
		if sueldo >= 100000.00 {
			fmt.Println("Se puede otorgar un credito sin intereses")
		} else {
			fmt.Println("Se puede otorgar un credito")
		}
	} else {
		fmt.Println("No se puede otorgar credito ya que no cumple con las reglas")
	}

	fmt.Println("=================")
}

func ejercicio4() {
	fmt.Println("Ejercicio 4")
	// se puede resolver con una concatenacion de if o switch
	switch numeroMes := 11; numeroMes {
	case 1:
		fmt.Println(numeroMes, "Enero")
	case 2:
		fmt.Println(numeroMes, "Febrero")
	case 3:
		fmt.Println(numeroMes, "Marzo")
	case 4:
		fmt.Println(numeroMes, "Abril")
	case 5:
		fmt.Println(numeroMes, "Mayo")
	case 6:
		fmt.Println(numeroMes, "Junio")
	case 7:
		fmt.Println(numeroMes, "Julio")
	case 8:
		fmt.Println(numeroMes, "Agosto")
	case 9:
		fmt.Println(numeroMes, "Septiembre")
	case 10:
		fmt.Println(numeroMes, "Octubre")
	case 11:
		fmt.Println(numeroMes, "Noviembre")
	case 12:
		fmt.Println(numeroMes, "Diciembre")
	default:
		fmt.Println(numeroMes, "No corresponde a un mes, intenta del 1 al 12")
	}

	fmt.Println("=================")
}

func ejercicio5() {
	fmt.Println("Ejercicio 5")
	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(estudiantes)

	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println("Estudiantes luego de 2 clases:")
	fmt.Println(estudiantes)
	fmt.Println("=================")
}

func ejercicio6() {
	fmt.Println("Ejercicio 6")
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	empleadoInteresado := "Benjamin" //esta variable se utiliza para buscar al empleado que se requiera informacion
	fmt.Println("La edad de", empleadoInteresado, "es:", empleados[empleadoInteresado])

	fmt.Println("Empleados mayores a 21")
	for key, element := range empleados {
		if element > 21 {
			fmt.Println("Nombre:", key, "-", "Edad:", element)
		}
	}

	empleados["Federico"] = 25
	fmt.Println("Se agregó Federico a la lista")

	delete(empleados, "Pedro")
	fmt.Println("Se eliminó a Pedro del mapa")

	fmt.Println("=================")
}

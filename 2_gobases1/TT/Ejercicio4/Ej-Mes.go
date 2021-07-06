package main

import "fmt"

func main() {
	var mes uint16 = 1

	fmt.Println("Ingrese el numero de un mes: ")
	fmt.Scanln(&mes)

	switch mes {
	case 1:
		fmt.Println("Mes ", mes, "- Enero.")
	case 2:
		fmt.Println("Mes ", mes, "- Febrero.")
	case 3:
		fmt.Println("Mes ", mes, "- Marzo.")
	case 4:
		fmt.Println("Mes ", mes, "- Abril.")
	case 5:
		fmt.Println("Mes ", mes, "- Mayo.")
	case 6:
		fmt.Println("Mes ", mes, "- Junio.")
	case 7:
		fmt.Println("Mes ", mes, "- Julio.")
	case 8:
		fmt.Println("Mes ", mes, "- Agosto.")
	case 9:
		fmt.Println("Mes ", mes, "- Septiembre.")
	case 10:
		fmt.Println("Mes ", mes, "- Octubre.")
	case 11:
		fmt.Println("Mes ", mes, "- Noviembre.")
	case 12:
		fmt.Println("Mes ", mes, "- Diciembre.")

	}
}

package main

import (
	"fmt"
)

func main() {
	cliente := make([]map[string]int, 5)
	cliente[0] = make(map[string]int, 1)
	cliente[0]["edad"] = 23
	cliente[0]["sueldo"] = 110000
	cliente[0]["antiguedad"] = 13
	cliente[0]["Empleado"] = 1
	cliente[1] = make(map[string]int, 1)
	cliente[1]["edad"] = 24
	cliente[1]["sueldo"] = 90000
	cliente[1]["antiguedad"] = 13
	cliente[1]["Empleado"] = 1
	cliente[2] = make(map[string]int, 1)
	cliente[2]["edad"] = 21
	cliente[2]["sueldo"] = 110000
	cliente[2]["antiguedad"] = 13
	cliente[2]["Empleado"] = 1
	cliente[3] = make(map[string]int, 1)
	cliente[3]["edad"] = 24
	cliente[3]["sueldo"] = 90000
	cliente[3]["antiguedad"] = 11
	cliente[3]["Empleado"] = 1
	cliente[4] = make(map[string]int, 1)
	cliente[4]["edad"] = 24
	cliente[4]["sueldo"] = 90000
	cliente[4]["antiguedad"] = 13
	cliente[4]["Empleado"] = 0

	for _, element := range cliente {

		if element["edad"] > 22 && element["sueldo"] > 100000 && element["antiguedad"] > 12 && element["Empleado"] == 1 {
			fmt.Println("el empleado Camilo tiene beneficio con el banco por lo que se le puede hacer un prestamos y no se le cobran intereses")

		} else if element["edad"] > 22 && element["sueldo"] <= 100000 && element["antiguedad"] > 12 && element["Empleado"] == 1 {
			fmt.Println("Al empleado Lukas se le hace el prestamo del banco pero se le cobran intereses")
		} else if element["edad"] <= 22 && element["sueldo"] > 100000 && element["antiguedad"] > 12 && element["Empleado"] == 1 {
			fmt.Println("Al empleado Felipe no se le puede hacer el prestamo por su edad, ya que no es mayor de 22 años")
		} else if element["edad"] > 22 && element["sueldo"] <= 100000 && element["antiguedad"] <= 12 && element["Empleado"] == 1 {
			fmt.Println("Al empleado Andres no se le puede realizar el prestamo porque no cumple con la antiguedad minima necesaria")
		} else if element["edad"] > 22 && element["sueldo"] <= 100000 && element["antiguedad"] > 12 && element["Empleado"] == 0 {
			fmt.Println("Al señor Luis no se le puede realizar un prestmo porque no es empleado de la empresa ")

		}

	}

}

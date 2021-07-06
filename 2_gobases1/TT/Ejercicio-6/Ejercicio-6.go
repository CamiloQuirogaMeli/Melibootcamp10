package main

import "fmt"

func printMap(MapParameter map[string]int) {
	for k, v := range MapParameter {
		fmt.Println("Nombre:", k, "\n\tEdad", v)
	}
}

func main() {
	fmt.Println("Ejercicio 6")
	count := 0
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])
	fmt.Println("Los empleados con edad mayor a 21 años son:")
	for key, value := range employees {
		if value > 21 {
			fmt.Println("-", key)
			count++
		}
	}
	fmt.Println("La cantidad de empleados mayor a 21 años son:", count)
	fmt.Println("Se agregara un nuevo empleado, la lista actual es:")
	printMap(employees)
	delete(employees, "Pedro")
	fmt.Println("Luego de la eliminacion de Pedro:")
	printMap(employees)
}

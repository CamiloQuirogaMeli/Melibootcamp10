package main

import "fmt"

func main() {
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int
	fmt.Println("Benjamin tiene ", empleados["Benjamin"], " años")

	for _, edad := range empleados {
		if edad >= 21 {
			mayores++
		}
	}
	fmt.Println("Existen ", mayores, " empleados mayores a 21 años")

	empleados["Federico"] = 25
	fmt.Println(empleados)

	delete(empleados, "Pedro")
	fmt.Println(empleados)
}

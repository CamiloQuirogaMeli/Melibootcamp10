package main

import "fmt"

func main() {
	var mayores21 int = 0
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19,
		"Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es: ", empleados["Benjamin"])

	for _, empleado := range empleados {
		if empleado > 21 {
			mayores21++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21: ", mayores21)

	fmt.Print(empleados)
	fmt.Println("")
	empleados["Federico"] = 25
	fmt.Println("Agrego a Federico")
	delete(empleados, "Pedro")
	fmt.Println("Elimino a Pedro")

	fmt.Print(empleados)
}

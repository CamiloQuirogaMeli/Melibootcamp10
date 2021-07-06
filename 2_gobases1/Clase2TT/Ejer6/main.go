package main

import (
	"fmt"
)

func main() {
	employees := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	overTwentyOne := 0
	for _, value := range employees {
		if value > 21 {
			overTwentyOne += 1
		}
	}
	fmt.Printf("El numero de empleados mayores de 21 años son %d\n", overTwentyOne)
	fmt.Println(employees)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}

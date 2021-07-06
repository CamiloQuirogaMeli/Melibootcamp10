package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Benjamin age %d\n\n", employees["Benjamin"])

	delete(employees, "Pedro")

	employees["Federico"] = 25

	for name, age := range employees {

		if age >= 21 {
			fmt.Printf("%s is 21 old \n", name)
		}
	}
}

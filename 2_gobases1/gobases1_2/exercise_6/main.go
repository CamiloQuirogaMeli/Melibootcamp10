package main

import "fmt"

var (
	employees     = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	c         int = 0
)

func main() {

	fmt.Println("Benjamin's age is", employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			c++
		}
	}

	fmt.Println("Employees over 21 years of age: ", c)

	employees["Federico"] = 25
	fmt.Println("Employees: ", employees)

	delete(employees, "Pedro")
	fmt.Println("Employees: ", employees)

}

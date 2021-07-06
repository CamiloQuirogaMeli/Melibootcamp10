package main

import "fmt"

// countEmployeesByAge count the employees older than 22 years old
func countEmployeesByAge(employees map[string]int) {
	var olderThan22 int = 0
	for _, age := range employees {
		if age > 21 {
			olderThan22++
		}
	}
	fmt.Printf("There are %d older than 22", olderThan22)
}

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad de benjamin: ", employees["Benjamin"])

	countEmployeesByAge(employees)

	fmt.Println("Adding employee Federico and deleting Pedro...")
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Updated employees list:", employees)
}

package main

import (
	f "fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0
	f.Println("Benjamin has", employees["Benjamin"], "years old")
	for _, value := range employees {
		if value > 21 {
			count++
		}
	}

	f.Println("Employeers that have more than 21 years old:", count)
	employees["Federico"] = 25
	delete(employees, "Pedro")

	f.Println("Employeers list updated:")
	count = 0
	for key, value := range employees {
		count++
		f.Printf("%d.- %s has %d years old\n", count, key, value)
	}
}
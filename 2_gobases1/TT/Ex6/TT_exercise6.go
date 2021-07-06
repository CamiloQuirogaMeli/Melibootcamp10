package main

import "fmt"

func main() {

	var count int = 0

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin age:", employees["Benjamin"])

	for _, employe := range employees {

		if employe > 21 {
			count++
		}

	}
	fmt.Println("employees over 21:", count)

}

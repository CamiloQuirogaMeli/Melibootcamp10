package main

import "fmt"

func main() {
	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}

	for key, element := range students {
		fmt.Println("Key: ", key, "=>", "Element: ", element)
	}
}

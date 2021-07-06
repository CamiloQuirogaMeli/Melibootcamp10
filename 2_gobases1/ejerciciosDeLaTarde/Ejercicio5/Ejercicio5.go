package main

import "fmt"

func main() {

	mySlice := []string{"Benjamin", "Nahuel", "Brenda", "Marcos"}

	fmt.Println(mySlice[:])

	var otherName string = "Pepe"

	mySlice = append(mySlice, otherName)

	fmt.Println(mySlice[:])
}

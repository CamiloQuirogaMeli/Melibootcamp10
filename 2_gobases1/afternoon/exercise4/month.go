package main

import "fmt"

func main() {
	month := 8

	numToMonth := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	if month < 1 || month > 12 {
		fmt.Println(month, "is not a valid month")
	} else {
		fmt.Println(month, "corresponds to", numToMonth[month-1])
	}

	//I could have used a series of consecutive if-else statements, or use a switch, but I think using an array is a less verbose solution, even, may a say, cleaner solution.
}

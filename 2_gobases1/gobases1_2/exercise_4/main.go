package main

import "fmt"

var (
	month_number uint8
)

func main() {

	fmt.Print("Enter the number of the month: ")
	fmt.Scanln(&month_number)

	// Alternative 1: Arrays

	month := [12]string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	fmt.Println(month_number, month[month_number-1])

	// Alternative 2: Switch

	switch month_number {
	case 1:
		fmt.Println(month_number, "January")
	case 2:
		fmt.Println(month_number, "February")
	case 3:
		fmt.Println(month_number, "March")
	case 4:
		fmt.Println(month_number, "April")
	case 5:
		fmt.Println(month_number, "May")
	case 6:
		fmt.Println(month_number, "June")
	case 7:
		fmt.Println(month_number, "July")
	case 8:
		fmt.Println(month_number, "August")
	case 9:
		fmt.Println(month_number, "September")
	case 10:
		fmt.Println(month_number, "October")
	case 11:
		fmt.Println(month_number, "November")
	case 12:
		fmt.Println(month_number, "December")
	}
}

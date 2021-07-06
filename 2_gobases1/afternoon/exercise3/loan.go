package main

import "fmt"

func main() {

	age := 25
	isEmployed := false
	yearsInJob := 2
	salary := 20000

	if age > 22 && isEmployed && yearsInJob > 1 {
		//loan is going to be aproved, lets check is client hast to pay any interest rate
		if salary >= 100000 {
			fmt.Println("Congrats! Your loan has been approved and you don't have to pay any interest rate")
		} else {
			fmt.Println("Congrats! Your loan has been approved")
		}
	} else {
		fmt.Println("We are sorry to inform you that you are not elegible for a loan right now")
	}
}

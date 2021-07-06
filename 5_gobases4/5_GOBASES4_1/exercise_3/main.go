package main

import (
	"fmt"
)

func verificationMessageTest(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: the taxable minimum is 150.000 and the salary entered is: %d", salary)
	}
	return nil
}

func main() {
	var salary int

	fmt.Println("Enter salary: ")
	fmt.Scanln(&salary)

	err := verificationMessageTest(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}

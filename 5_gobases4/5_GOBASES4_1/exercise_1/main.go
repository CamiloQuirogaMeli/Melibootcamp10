package main

import "fmt"

type verificationMessage struct {
	msg string
}

func (v *verificationMessage) Error() string {
	return fmt.Sprint(v.msg)
}

func verificationMessageTest(salary int) error {
	if salary < 150000 {
		return &verificationMessage{
			msg: "error: the salary entered does not reach the taxable minimum",
		}
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

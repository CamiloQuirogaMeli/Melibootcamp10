package main

import "fmt"

func main() {
	age := 24
	emp := false
	sen := 16

	req1, req2, req3 := loan(age, emp, sen)

	if req1 && req2 && req3 {
		fmt.Println("All requirements accomplished. Your loan is Approved")
	} else {
		fmt.Println("One or more requirements are not accomplished. Loan is denied :(")
	}
}

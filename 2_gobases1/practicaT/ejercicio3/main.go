package main

import "fmt"

func main() {

	var userName string = "test"
	type User struct {
		age      int
		hasJob   bool
		workTime int
		salary   int
	}

	users := make(map[string]User)
	users[userName] = User{23, true, 10, 1000000000}

	var user User = users[userName]
	if user.age > 22 && user.hasJob && user.workTime > 1 {
		if user.salary > 100000 {
			fmt.Println("Accede a prestamos sin cobro de interes")
		} else {
			fmt.Println("Accede a prestamos con cobro de interes")
		}

	} else {
		fmt.Println("No puede acceder a ningun prestamos")
	}
}

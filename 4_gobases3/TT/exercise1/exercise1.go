package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (u *User) setName(name *string) {
	u.name = *name
}

func (u *User) setLastName(lastName *string) {
	u.lastName = *lastName
}

func (u *User) setAge(age *int) {
	u.age = *age
}

func (u *User) setEmail(email *string) {
	u.email = *email
}

func (u *User) setPassword(password *string) {
	u.password = *password
}

func main() {
	user := User{}
	var (
		name     = "Carlos"
		lastName = "Infante"
		age      = 20
		email    = "carlos.infante@mercadolibre.com.co"
		password = "abc12345"
	)
	user.setName(&name)
	user.setLastName(&lastName)
	user.setAge(&age)
	user.setEmail(&email)
	user.setPassword(&password)
	fmt.Println(user)
}

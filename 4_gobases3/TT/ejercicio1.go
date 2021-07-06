package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int8
	email    string
	password string
}

func (u *User) setName(name string, lastName string) {
	if name != "" {
		u.name = name
	}
	if lastName != "" {
		u.lastName = lastName
	}
}

func (u *User) setAge(age int8) {
	u.age = age
}

func (u *User) setEmail(email string) {
	u.email = email
}

func (u *User) setPassword(pass string) {
	u.password = pass
}

func main() {
	var user = User{}
	user.setName("Alejandro", "Castillo")
	user.setAge(25)
	user.setEmail("alejandro.castillo@mercadolibre.com.mx")
	user.setPassword("123")
	fmt.Println(user)
}

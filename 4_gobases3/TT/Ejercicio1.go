package main

import "fmt"

type user struct {
	firstName string
	lastName string
	age int
	mail string
	password string
}

func main() {
	user := user{
		firstName: "Pepe",
		lastName: "Toto",
		age: 55,
		mail: "pepe@ml.com",
		password: "123",
	}
	fmt.Println(user)
	user.SetAge(34)
	user.SetFirstName("lolo")
	fmt.Println(user)
}

func (user *user) SetFirstName(name string) {
	user.firstName = name
}
func (user *user) SetLastName(lastName string) {
	user.lastName = lastName
}
func (user *user) SetAge(age int) {
	user.age = age
}
func (user *user) SetMail(mail string) {
	user.mail = mail
}
func (user *user) SetPassword(password string) {
	user.password = password
}


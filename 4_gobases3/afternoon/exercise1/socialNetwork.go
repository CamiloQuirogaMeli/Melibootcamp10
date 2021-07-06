package main

import "fmt"

type User struct {
	name     string
	age      int
	email    string
	password string
}

func main() {

	user := User{"Andres", 24, "afmoyar", "secret"}
	fmt.Println("Original user: ", user)
	changeUser(&user)
	fmt.Println("User after update:", user)

}

func changeUser(user *User) {
	user.setName("Felipe")
	user.setAge(21)
	user.setEmail("anmoyarodrig")
	user.setPassword("super secret")
}

func (user *User) setName(theName string) {
	user.name = theName
}
func (user *User) setAge(theAge int) {
	user.age = theAge
}
func (user *User) setEmail(theEmail string) {
	user.email = theEmail
}
func (user *User) setPassword(thePassword string) {
	user.password = thePassword
}

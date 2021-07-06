package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func changeName(user *User, newName string) {
	user.name = newName
}

func changeAge(user *User, newAge int) {
	user.age = newAge
}

func changeEmail(user *User, newEmail string) {
	user.email = newEmail
}

func changePassword(user *User, newPassword string) {
	user.password = newPassword
}

func main() {

	user := User{"Gianluca", "Ciccarelli", 20, "gian.cicca@gmail.com", "pwdfake1a2s3d"}

	changeName(&user, "Gian")
	changeAge(&user, 21)
	changeEmail(&user, "gianluca.ciccarelli@mercadolibre.com")
	changePassword(&user, "anotherfakepwd1a2sd213")

	fmt.Print(user)
}

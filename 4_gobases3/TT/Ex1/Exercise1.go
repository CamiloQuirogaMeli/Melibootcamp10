package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	age       int
	mail      string
	password  string
}

func (u *user) changeName(firstName *string, lastName *string) {

	u.firstName = *firstName
	u.lastName = *lastName

}

func (u *user) changeAge(age *int) {

	u.age = *age

}
func (u *user) changeMail(mail *string) {

	u.mail = *mail

}
func (u *user) changePassword(password *string) {

	u.password = *password

}

func main() {

	user := user{"John", "Doe", 25, "john.doe@asd.com", "1234"}

	newFirstName := "Jane"
	newLastName := "Does"
	newAge := 30
	newMail := "jane.does@asd.com"
	newPassword := "9876"

	user.changeName(&newFirstName, &newLastName)
	user.changeAge(&newAge)
	user.changeMail(&newMail)
	user.changePassword(&newPassword)

	fmt.Println(user)
}

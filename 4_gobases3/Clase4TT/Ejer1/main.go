package main

import (
	"fmt"
)

type User struct {
	name     string
	lastName string
	age      int64
	email    string
	pwd      string
}

func (u *User) changeName(name, lastName *string) {
	(*u).name = *name
	(*u).lastName = *lastName
}

func (u *User) changeAge(age *int64) {
	(*u).age = *age
}

func (u *User) changeEmail(email *string) {
	(*u).email = *email
}

func (u *User) changePwd(pwd *string) {
	(*u).pwd = *pwd
}

func main() {
	user := User{"Will", "Garcia", 25, "test@test.com", "123456"}
	fmt.Println(user)
	newName := "Tom"
	newLastname := "sawyer"
	newPwd := "654321"
	newEmail := "newTest@newTest.com"
	var newAge int64 = 12
	user.changeName(&newName, &newLastname)
	user.changeAge(&newAge)
	user.changeEmail(&newEmail)
	user.changePwd(&newPwd)
	fmt.Println(user)
}

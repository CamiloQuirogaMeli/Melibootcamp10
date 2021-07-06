package main

import (
	"fmt"
)

type Student struct {
	name     string
	surname  string
	age      int
	email    string
	password string
}

func (student *Student) setName(name string, surname string) {
	student.name = name
	student.surname = surname
}

func (student *Student) setAge(age int) {
	student.age = age
}

func (student *Student) setEmail(email string) {
	student.email = email
}

func (student *Student) setPassword(password string) {
	student.password = password
}
func main() {

	x := Student{}
	x.setName("Horacio", "Perez")
	x.setAge(22)
	x.setEmail("hora@gmail.com")
	x.setPassword("xxxxxxxx")
	fmt.Println(x)

}

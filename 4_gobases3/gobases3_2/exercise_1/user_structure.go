package main

import f "fmt"

type User struct {
	name     string
	surname  string
	age      int
	email    string
	password string
}

func (u User) printData() {
	f.Println("\nName:", u.name, u.surname)
	f.Println("Age:", u.age)
	f.Println("Email:", u.email)
	f.Println("Password:", u.password)
	f.Println("Password successfully changed")
}

package main

import (
	f "fmt"
)

var (
	new_name     string
	new_surname  string
	new_age      int
	new_email    string
	new_password string
)

func changeName(u *User) {
	f.Print("\nEnter the new name: ")
	f.Scan(&new_name)
	f.Print("Enter the new surname: ")
	f.Scan(&new_surname)
	u.name = new_name
	u.surname = new_surname
	f.Println("Name successfully changed")
}

func changeAge(u *User) {
	f.Print("\nEnter the new age: ")
	f.Scan(&new_age)
	u.age = new_age
	f.Println("Age successfully changed")
}

func changeEmail(u *User) {
	f.Print("\nEnter the new email: ")
	f.Scan(&new_email)
	u.email = new_email
	f.Println("Email successfully changed")
}

func changePassword(u *User) {
	f.Print("\nEnter the new password: ")
	f.Scan(&new_password)
	u.password = new_password
	f.Println("Password successfully changed")
}

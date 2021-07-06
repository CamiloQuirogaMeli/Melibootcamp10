package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type User struct {
	name     string
	age      int
	email    string
	password string
}

func (u *User) SetName(n string) error {
	if n == "" {
		return errors.New("no name provided")
	}
	u.name = n
	return nil
}

func (u *User) SetAge(age int) error {
	if age < 1 || age > 100 {
		return errors.New("invalid age")
	}
	u.age = age
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("no email provided")
	}
	u.email = email
	return nil
}
func (u *User) SetPassword(pw string) error {
	if pw == "" {
		return errors.New("no password provided")
	}
	u.password = pw
	return nil
}

func (u *User) Info() {
	fmt.Printf("Name: \t %s\nAge: \t %d \nEmail: \t %s\nPassw: \t %s\n", u.name, u.age, u.email, u.password)
}
func main() {
	user := &User{}
	var err error
	err = user.SetName("Marcos")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	err = user.SetAge(24)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	err = user.SetEmail("marcos.zabala@mercadolibre.com")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	err = user.SetPassword("this is a pw")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	user.Info()
}

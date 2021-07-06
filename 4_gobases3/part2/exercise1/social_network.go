package main

import (
	f "fmt"
)

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

func (user *User) changeName(name string) {
	user.Name = name
}

func (user *User) changeLastName(lastname string) {
	user.Lastname = lastname
}

func (user *User) changeAge(age int) {
	user.Age = age
}

func (user *User) changeEmail(email string) {
	user.Email = email
}

func (user *User) changePassword(password string) {
	user.Password = password
}

func main() {
	usuario := User{Name: "Ale", Lastname: "lamar", Age: 23, Email: "ale.jaam", Password: "hola"}
	f.Println("Informacion del usaurio:")
	f.Println("Name:", usuario.Name)
	f.Println("Lastname:", usuario.Lastname)
	f.Println("Age:", usuario.Age)
	f.Println("Email:", usuario.Email)
	f.Println("Password:", usuario.Password)
	usuario.changeName("Alejandro")
	usuario.changeLastName("Alamar")
	usuario.changeAge(24)
	usuario.changeEmail("ale.jaam7@gmail.com")
	usuario.changePassword("Hola123")
	f.Println("-------------------------------------")
	f.Println("Informacion del usaurio modificada:")
	f.Println("Name:", usuario.Name)
	f.Println("Lastname:", usuario.Lastname)
	f.Println("Age:", usuario.Age)
	f.Println("Email:", usuario.Email)
	f.Println("Password:", usuario.Password)
}

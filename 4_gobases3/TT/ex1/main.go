package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func (u *User) setName(n string) *User {
	u.Name = n
	return u
}

func (u *User) setAge(a string) *User {
	u.Age, _ = strconv.Atoi(a)
	return u
}

func (u *User) setEmail(e string) *User {
	u.Email = e
	return u
}

func (u *User) setPassword(p string) *User {
	u.Password = p
	return u
}

func main() {
	user := User{
		Name:     "John Doe",
		Age:      21,
		Email:    "johndoe@example.com",
		Password: "123456",
	}
	for true {
		fmt.Println("Usuario:")
		fmt.Println("Nombre:     ", user.Name)
		fmt.Println("Edad:       ", user.Age)
		fmt.Println("Email:      ", user.Email)
		fmt.Println("Contraseña: ", user.Password)
		fmt.Println()
		menu := ""
		fmt.Print("¿Desea cambiar algún dato? (S/N): ")
		fmt.Scanln(&menu)
		menu = strings.ToLower(menu)
		if menu == "s" {
			fmt.Println("¿Que desea cambiar?")
			fmt.Println(" * (1) Nombre")
			fmt.Println(" * (2) Edad")
			fmt.Println(" * (3) Email")
			fmt.Println(" * (4) Contraseña")
			fmt.Scanln(&menu)
			switch menu {
			case "1":
				n := ""
				fmt.Print("Nuevo nombre: ")
				fmt.Scanln(&menu)
				n += menu + " "
				fmt.Print("Nuevo apellido: ")
				fmt.Scanln(&menu)
				n += menu
				user.setName(n)
			case "2":
				fmt.Print("Nueva edad: ")
				fmt.Scanln(&menu)
				user.setAge(menu)
			case "3":
				fmt.Print("Nuevo email: ")
				fmt.Scanln(&menu)
				user.setEmail(menu)
			case "4":
				fmt.Print("Nueva contraseña: ")
				fmt.Scanln(&menu)
				user.setPassword(menu)
			}
		} else if menu == "N" {
			return
		}
	}
}

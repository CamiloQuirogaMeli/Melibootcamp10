package main

import "fmt"

func main() {
	var pointerUser *User
	userA := User{Name: "Sebastian", LastName: "Posadas", Email: "sebastian.posadas@mercadolibre.com", Password: "123456"}

	fmt.Println(userA.ToString())

	pointerUser = &userA

	setFullName(pointerUser, "Simon", "Pedrito")
	setEmail(pointerUser, "simon.posadas@mercadolibre.com")
	setPassword(pointerUser, "1234")
	setAge(pointerUser, 2)

	fmt.Println(userA.ToString())
}

type User struct {
	Name     string
	LastName string
	Email    string
	Age      int
	Password string
}

func setFullName(u *User, name, lastName string) {
	(*u).Name = name
	(*u).LastName = lastName
}

func setEmail(u *User, email string) {
	(*u).Email = email
}

func setAge(u *User, age int) {
	(*u).Age = age
}

func setPassword(u *User, password string) {
	(*u).Password = password
}

func (u User) ToString() string {
	return fmt.Sprintf("name: %s, lastName: %s, email: %s, age: %v, password: %s", u.Name, u.LastName, u.Email, u.Age, u.Password)
}

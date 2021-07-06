package main

func main() {
	var user_pointer *User
	user := User{
		name:     "Joseph",
		surname:  "Caicedo",
		age:      20,
		email:    "joseph.saenz@mercadolibre.com.co",
		password: "kjTy12314.",
	}
	user_pointer = &user

	user.printData()

	changeName(user_pointer)

	changeAge(user_pointer)

	changeEmail(user_pointer)

	changePassword(user_pointer)

	user.printData()
}

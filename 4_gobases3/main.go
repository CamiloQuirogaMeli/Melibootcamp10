package main

func main() {
	/* 	saveFile()
	   	readFile() */

	var (
		nombre     = "Vincent"
		apellido   = "Ramon"
		edad       = 23
		correo     = "vincentconace@gmail.com"
		contrasena = "12345435"
	)

	user1 := User{nombre, apellido, edad, correo, contrasena}
	user1.print()
	user1.changeName("Carlos", "Gonzalez")
	user1.changeAge(26)
	user1.changeMail("carlos@gmail.com")
	user1.changePass("1234567890")
	user1.print()
}

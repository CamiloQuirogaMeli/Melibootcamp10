package main

import "fmt"
import "encoding/json"

func main() {
	
	fmt.Println("* Caso funciones de User modifican sus datos")
	var userJson []byte
	user := User{
		Name : "Homero",
		Surname : "Simpson",
		Age : 50,
		Mail : "homero-simpson@springfield.com",
		Password : "Calle Falsa 123",
	}
	userJson, _ = json.MarshalIndent(user, " ", " ")
	fmt.Println("Usuario antes de los cambios:\n", string(userJson))
	user.
		ChangeNameSurname("Sponge", "Bob").
		ChangeMail("spongebob@bikini.com").
		ChangeAge(12).
		ChangePassword("Patrick")
	userJson, _ = json.MarshalIndent(user, "", " ")
	fmt.Println("Usuario despues de los cambios:\n", string(userJson))
}

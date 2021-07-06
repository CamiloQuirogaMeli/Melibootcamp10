package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type user struct {
	Id        int
	FirstName string
	LastName  string
	Mail      string
	Age       int
	Height    float64
	Active    bool
	Created   string
}

func main() {

	user1 := user{1, "John", "Doe", "johndoe@asd.com", 24, 1.75, true, "15/04/2020"}
	user2 := user{2, "Jane", "Does", "janedoes@asd.com", 40, 1.80, true, "6/12/2019"}

	users := []user{user1, user2}

	json, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Error:", err)
	} else {

		err = ioutil.WriteFile("users.json", []byte(json), 0644)

		if err != nil {
			fmt.Println("Error:", err)

		} else {
			fmt.Println("Success!")

		}

	}

}

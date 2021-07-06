package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Age       int
	Height    float64
	Active    bool
	Created   string
}

func main() {

	usersJson, err := ioutil.ReadFile("../../users.json")

	if err != nil {
		fmt.Println("Error reading file")
	}

	var users []user

	json.Unmarshal(usersJson, &users)

	router := gin.Default()

	router.GET("/users/getAll", func(c *gin.Context) {

		c.JSON(200, users)

	})

	router.Run()
}

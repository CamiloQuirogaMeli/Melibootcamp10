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

var users []user

func main() {

	usersJson, err := ioutil.ReadFile("../../users.json")

	if err != nil {
		fmt.Println("Error reading file")
	}

	json.Unmarshal(usersJson, &users)

	router := gin.Default()

	usersGroup := router.Group("/users")
	{
		usersGroup.POST("", addUser())
	}
	router.Run()
}

func addUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user user

		if err := c.Bind(&user); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastId := users[len(users)-1].Id

		user.Id = lastId + 1

		users = append(users, user)

		json, _ := json.Marshal(users)

		ioutil.WriteFile("../../users.json", []byte(json), 0644)

		c.JSON(200, user)

	}
}

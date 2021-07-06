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

		if user.FirstName == "" {
			c.String(400, "the field first_name is required")
			return
		}
		if user.LastName == "" {
			c.String(400, "the field last_name is required")
			return
		}
		if user.Email == "" {
			c.String(400, "the field email is required")
			return
		}
		if user.Age == 0 {
			c.String(400, "the field age is required")
			return
		}
		if user.Height == 0.0 {
			c.String(400, "the field height is required")
			return
		}
		if user.Created == "" {
			c.String(400, "the field created is required")
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

package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

type User struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	LastName    string  `json:"lastName`
	Email       string  `json:"email"`
	Age         int     `json:"age"`
	Height      float64 `json:"height"`
	Active      bool    `json:"active"`
	DateCreated string  `json:"dateCreated"`
}

func GetAllUsers(c *gin.Context) {

	file, _ := ioutil.ReadFile("./users.json")

	users := []User{}

	_ = json.Unmarshal(file, &users)

	var usersFiltered []User
	for _, user := range users {

		status, _ := strconv.ParseBool(c.Query("active"))
		if status == user.Active {
			usersFiltered = append(usersFiltered, user)
		}
	}

	if len(usersFiltered) != 0 {

		c.JSON(200, gin.H{
			"message": usersFiltered,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": users,
	})
}

func GetOneUser(c *gin.Context) {

	file, _ := ioutil.ReadFile("./users.json")

	users := []User{}

	_ = json.Unmarshal(file, &users)

	var userFiltered User
	userId, _ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.Id == userId {
			userFiltered = user
		}
	}

	if userFiltered == (User{}) {

		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": userFiltered,
	})
}

func main() {

	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.GET("/users", GetAllUsers)

	router.GET("/user/:id", GetOneUser)

	router.Run()
}

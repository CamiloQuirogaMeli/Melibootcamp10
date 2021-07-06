package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name        string  `json:"name" binding:"required"`
	LastName    string  `json:"lastName" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	Age         int     `json:"age" binding:"required"`
	Height      float64 `json:"height" binding:"required"`
	Active      bool    `json:"active" binding:"required"`
	DateCreated string  `json:"dateCreated" binding:"required"`
}

type CreatedUser struct {
	Id int `json:"id"`
	User
}

var lastId int

const TOKEN = "1234"

func CreateUsers(c *gin.Context) {

	token := c.Request.Header.Get("token")

	if token != TOKEN {
		c.JSON(401, "you do not have access to make the requested petition")
		return
	}

	var req []User
	if err := c.Bind(&req); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUsers := make([]CreatedUser, len(req))

	for i, user := range req {

		createdUsers[i] = CreatedUser{
			Id:   lastId,
			User: user,
		}
		lastId++
	}

	c.JSON(201, gin.H{
		"createdUsers": createdUsers,
	})
}

func main() {

	router := gin.Default()

	router.POST("/user", CreateUsers)

	router.Run()
}

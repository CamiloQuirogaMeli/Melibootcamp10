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
		usersGroup.GET("/:id", getUser)
	}
	router.Run()
}

func getUser(ctx *gin.Context) {

	filId := ctx.Param("id")

	for _, v := range users {

		if fmt.Sprintf("%v", v.Id) == filId {
			ctx.JSON(200, v)
			return
		}
	}

	ctx.String(404, "User not found")

}

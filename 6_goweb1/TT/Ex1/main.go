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
		usersGroup.GET("", getUser)
	}
	router.Run()
}

func getUser(ctx *gin.Context) {

	res := []user{}

	filFirstName := ctx.Query("first_name")
	filLastName := ctx.Query("last_name")
	filEmail := ctx.Query("email")
	filAge := ctx.Query("age")
	filHeight := ctx.Query("height")
	filActive := ctx.Query("active")
	filCreated := ctx.Query("created")

	for _, v := range users {

		if (filFirstName == "" || v.FirstName == filFirstName) && (filLastName == "" || v.LastName == filLastName) && (filEmail == "" || v.Email == filEmail) && (filAge == "" || fmt.Sprintf("%v", v.Age) == filAge) && (filHeight == "" || fmt.Sprintf("%.2f", v.Height) == filHeight) && (filActive == "" || v.Active == (filActive == "1")) && (filCreated == "" || v.Created == filCreated) {
			res = append(res, v)
		}
	}

	ctx.JSON(200, res)

}

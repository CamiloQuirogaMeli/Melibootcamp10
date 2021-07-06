package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	//EJERCICIO 3
	router := gin.Default()
	router.GET("/users", GetAll)
	router.Run()

}

type User struct {
	Id           int
	Name         string
	Surname      string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate string
}

func GetAll(c *gin.Context) {
	users := []User{}
	dJson, _ := ioutil.ReadFile("users.json")
	err := json.Unmarshal(dJson, &users)
	if err != nil {
		panic("json fail")
	}
	c.JSON(200, users)
}

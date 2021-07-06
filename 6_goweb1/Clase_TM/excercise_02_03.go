package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	name := "Sebastian"

	router := gin.Default()
	router.GET("/greetings", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	jsonUsers, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}
	if err := json.Unmarshal(jsonUsers, &users); err != nil {
		log.Fatal(err)
	}

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, users)
	})

	router.Run()
}

type User struct {
	Id           int
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       float64
	Active       bool
	CreationDate time.Time
}

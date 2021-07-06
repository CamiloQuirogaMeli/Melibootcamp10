package main

import (
	"encoding/json"
	"fmt"

	// "github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	data, err := ioutil.ReadFile("./users.json")

	if err != nil {
		fmt.Println(err)
	}

	users := []User{}
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
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

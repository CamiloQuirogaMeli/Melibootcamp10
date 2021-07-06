package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users/filters", filterUsers)
	router.GET("/users/:id", getUser)

	router.Run()
}

func filterUsers(ctx *gin.Context) {
	// var bindedUser User
	users := GetUsers()
	filteredUsers := []User{}

	// if ctx.Bind(&bindedUser) == nil {
	// 	fmt.Println(bindedUser)
	// }

	for _, user := range users {
		// if bindedUser.Id == user.Id || bindedUser.Name == user.Name || bindedUser.LastName == user.LastName || bindedUser.Email == user.Email ||
		// 	bindedUser.Age == user.Age || bindedUser.Height == user.Height || bindedUser.Active == user.Active || bindedUser.CreationDate == user.CreationDate {
		// 	filteredUsers = append(filteredUsers, user)
		// } //El problema con bindear, es que le da los valores por defecto a sus variables (int=0, bool=false) y chequear que sea == nil o != nill no funciona.
		if ctx.Query("id") == strconv.Itoa(user.Id) || ctx.Query("name") == user.Name ||
			ctx.Query("lastName") == user.LastName || ctx.Query("email") == user.Email ||
			ctx.Query("age") == strconv.Itoa(user.Age) || ctx.Query("height") == strconv.FormatFloat(user.Height, 'E', -1, 32) ||
			ctx.Query("active") == strconv.FormatBool(user.Active) || ctx.Query("creation_date") == user.CreationDate.String() {
			filteredUsers = append(filteredUsers, user)
		}
	}

	if len(filteredUsers) != 0 {
		// ctx.JSON(200, gin.H{
		// 	"users": filteredUsers,
		// })
		ctx.JSON(200, filteredUsers)
		return
	}

	ctx.JSON(200, users)
}

func getUser(ctx *gin.Context) {
	users := GetUsers()
	filteredUser := User{}

	for _, user := range users {
		if ctx.Param("id") == strconv.Itoa(user.Id) {
			filteredUser = user
			break
		}
	}

	if filteredUser == (User{}) {
		ctx.JSON(404, gin.H{
			"message": "User not found",
		})
	}

	ctx.JSON(200, filteredUser)

}

func GetUsers() []User {
	jsonUsers, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
	}

	users := []User{}
	if err := json.Unmarshal(jsonUsers, &users); err != nil {
		log.Fatal(err)
	}

	return users
}

type User struct {
	//TODO: Agregar el form para que sea posible el BIND
	Id           int       `form:"id" json:"id"`
	Name         string    `form:"name" json:"name"`
	LastName     string    `form:"lastName" json:"lastName"`
	Email        string    `form:"email" json:"email"`
	Age          int       `form:"age" json:"age"`
	Height       float64   `form:"height" json:"height"`
	Active       bool      `form:"active" json:"active"`
	CreationDate time.Time `form:"creation_date" json:"creation_date"`
}

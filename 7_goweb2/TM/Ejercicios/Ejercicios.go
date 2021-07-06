package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	/*
	   Ejercicios 1, 2 y 3
	*/
	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/", Save())
	r.Run()
}

func Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req User
		token := c.Request.Header.Get("token")
		if token != TOKEN {
			c.String(401, "usuario no autorizado")
			return
		}

		if err := c.Bind(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		validar := validacionDatos(req)

		if validar != "" {
			c.String(400, validar)
			return
		}

		req.Id = getById(users)
		req.Active = true
		users = append(users, req)
		c.JSON(200, req)

	}

}

var users = []User{}

const TOKEN = "1111"

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

func getById(u []User) int {
	var max int
	for _, v := range u {
		max = v.Id
	}
	max++
	return max
}

func validacionDatos(u User) string {
	switch {
	case u.Name == "":
		return "El campo nombre es requerido"
	case u.Surname == "":
		return "El campo apellido es requerido"
	case u.Email == "":
		return "El campo email es requerido"
	case u.Age == 0:
		return "El campo edad es requerido"
	case u.Height == 0:
		return "El campo altura es requerido"
	case u.CreationDate == "":
		return "El campo fecha creacion es requerido"
	default:
		return ""
	}
}

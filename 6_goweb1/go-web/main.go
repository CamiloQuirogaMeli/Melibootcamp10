package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id       int    `form:"id"`
	Nombre   string `form:"name"`
	Apellido string `form:"surname"`
	Email    string `form:"email"`
	Edad     int    `form:"age"`
	Altura   string `form:"height"`
	Activo   bool   `form:"active"`
	Fecha    string `form:"date"`
}

func GetAll(con *gin.Context) {
	var user []Usuario
	seFiltra := false
	dat, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		con.String(400, "ERROR")
	} else {

		err := json.Unmarshal(dat, &user)
		if err != nil {
			fmt.Println(err)
			con.String(400, "ERROR")
		} else {
			var filtrados []Usuario
			if con.Query("active") == "true" {
				seFiltra = true
				for _, value := range user {
					if value.Activo {
						filtrados = append(filtrados, value)
					}
				}
				if len(filtrados) > 0 {
					con.String(200, "Filtrado exitoso !", filtrados)
				} else {
					con.String(400, "No hubo resultados")
				}
			}
			if !seFiltra {
				con.String(200, "Sin filtro: ", user)
			}
		}
	}
}

func GetID(con *gin.Context) {
	var user []Usuario
	dat, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		con.String(400, "ERROR")
	} else {
		var usuario Usuario
		err := json.Unmarshal(dat, &user)
		if err != nil {
			fmt.Println(err)
			con.String(400, "ERROR")
		} else {
			id, err := strconv.Atoi(con.Param("id")[3:])
			if err != nil {
				fmt.Println(err)
			} else {
				for _, u := range user {
					if u.Id == id {
						usuario = u
					}
				}
				if usuario.Id > 0 {
					con.String(200, "Usuario encontrado: ", usuario)
				} else {
					con.String(404, "Usuario no encontrado")
				}
			}
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/usuarios", GetAll)
	router.GET("/usuarios/:id", GetID)
	router.Run()
}

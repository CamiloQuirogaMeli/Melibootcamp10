package main

import (
	"github.com/gin-gonic/gin"
)

type Silla struct {
	Id    string `json:"id"`
	Color string `json:"color"`
	Marca string `json:"marca"`
	Gamer bool   `json:"gamer"`
}

var arr = map[string]Silla{}

func getSilla(ctx *gin.Context) {

	silla, exisate := arr[ctx.Param("id")]

	if exisate {
		ctx.JSON(200, silla)
	} else {
		ctx.String(404, "no hay una silla con ese id")
	}

}

func main() {

	arr["1"] = Silla{"1", "rojo", "marca", false}
	arr["2"] = Silla{"2", "negro", "marca2", true}
	arr["3"] = Silla{"3", "blue", "marca3", false}

	router := gin.Default()

	router.GET("/silla/:id", getSilla)

	router.Run()
}

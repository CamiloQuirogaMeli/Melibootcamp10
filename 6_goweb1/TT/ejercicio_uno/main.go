package main

import (
	"github.com/gin-gonic/gin"
)

type Silla struct {
	Color string `json:"color"`
	Marca string `json:"marca"`
	Gamer bool   `json:"gamer"`
}

var arr = make([]Silla, 0)

func getByParam(ctx *gin.Context) {

	filtrados := make([]Silla, 0)
	cantFiltrados := 0

	for _, silla := range arr {

		if ctx.Query("gamer") != "" {
			buffGamer := false
			if ctx.Query("gamer") == "true" {
				buffGamer = true
			}

			if buffGamer != silla.Gamer {
				continue
			}

		}

		if ctx.Query("color") != "" {
			if ctx.Query("color") != silla.Color {
				continue
			}
		}

		if ctx.Query("marca") != "" {
			if ctx.Query("marca") != silla.Marca {
				continue
			}
		}

		filtrados = append(filtrados, silla)
		cantFiltrados++

	}

	if cantFiltrados > 0 {
		ctx.JSON(200, gin.H{
			"message": filtrados,
		})
	} else {
		ctx.String(404, "no hay Sillas que cumplan conn esa condicion")
	}

}

func main() {

	arr = append(arr, Silla{"rojo", "marca", false})
	arr = append(arr, Silla{"negro", "marca2", true})
	arr = append(arr, Silla{"azul", "marca3", false})

	router := gin.Default()

	router.GET("/sillas/", getByParam)

	router.Run()
}

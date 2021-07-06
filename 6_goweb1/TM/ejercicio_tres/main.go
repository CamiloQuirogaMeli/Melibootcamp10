package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Silla struct {
	Color string `json:"color"`
	Marca string `json:"marca"`
	Gamer bool   `json:"gamer"`
}

func getAll(ctx *gin.Context) {

	arr := make([]Silla, 0)
	arr = append(arr, Silla{"rojo", "marca", false})
	arr = append(arr, Silla{"negro", "marca2", true})
	arr = append(arr, Silla{"azul", "marca3", false})

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(arr)

	fmt.Println("->", *encoder)

	ctx.JSON(200, gin.H{
		"message": arr,
	})
}

func main() {
	router := gin.Default()

	router.GET("/sillas", getAll)

	router.Run()
}
